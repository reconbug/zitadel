package projection

import (
	"testing"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/handler"
	"github.com/caos/zitadel/internal/eventstore/repository"
	"github.com/caos/zitadel/internal/repository/iam"
)

func TestIAMMemberProjection_reduces(t *testing.T) {
	type args struct {
		event func(t *testing.T) eventstore.EventReader
	}
	tests := []struct {
		name   string
		args   args
		reduce func(event eventstore.EventReader) (*handler.Statement, error)
		want   wantReduce
	}{
		{
			name: "iam.MemberAddedType",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.MemberAddedEventType),
					iam.AggregateType,
					[]byte(`{
					"userId": "user-id",
					"roles": ["role"]
				}`),
				), iam.MemberAddedEventMapper),
			},
			reduce: (&IAMMemberProjection{}).reduceAdded,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IAMMemberProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "INSERT INTO zitadel.projections.iam_members (iam_id, user_id, roles, creation_date, change_date, sequence, resource_owner) VALUES ($1, $2, $3, $4, $5, $6, $7)",
							expectedArgs: []interface{}{
								"agg-id",
								"user-id",
								[]string{"role"},
								anyArg{},
								anyArg{},
								uint64(15),
								"ro-id",
							},
						},
					},
				},
			},
		},
		{
			name: "iam.MemberChangedType",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.MemberChangedEventType),
					iam.AggregateType,
					[]byte(`{
					"userId": "user-id",
					"roles": ["role", "changed"]
				}`),
				), iam.MemberChangedEventMapper),
			},
			reduce: (&IAMMemberProjection{}).reduceChanged,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IAMMemberProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "UPDATE zitadel.projections.iam_members SET (roles, change_date, sequence) = ($1, $2, $3) WHERE (iam_id = $4) AND (user_id = $5)",
							expectedArgs: []interface{}{
								[]string{"role", "changed"},
								anyArg{},
								uint64(15),
								"agg-id",
								"user-id",
							},
						},
					},
				},
			},
		},
		{
			name: "iam.MemberCascadeRemovedType",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.MemberCascadeRemovedEventType),
					iam.AggregateType,
					[]byte(`{
					"userId": "user-id"
				}`),
				), iam.MemberCascadeRemovedEventMapper),
			},
			reduce: (&IAMMemberProjection{}).reduceCascadeRemoved,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IAMMemberProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.iam_members WHERE (iam_id = $1) AND (user_id = $2)",
							expectedArgs: []interface{}{
								"agg-id",
								"user-id",
							},
						},
					},
				},
			},
		},
		{
			name: "iam.MemberRemovedType",
			args: args{
				event: getEvent(testEvent(
					repository.EventType(iam.MemberRemovedEventType),
					iam.AggregateType,
					[]byte(`{
					"userId": "user-id"
				}`),
				), iam.MemberRemovedEventMapper),
			},
			reduce: (&IAMMemberProjection{}).reduceRemoved,
			want: wantReduce{
				aggregateType:    iam.AggregateType,
				sequence:         15,
				previousSequence: 10,
				projection:       IAMMemberProjectionTable,
				executer: &testExecuter{
					executions: []execution{
						{
							expectedStmt: "DELETE FROM zitadel.projections.iam_members WHERE (iam_id = $1) AND (user_id = $2)",
							expectedArgs: []interface{}{
								"agg-id",
								"user-id",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := baseEvent(t)
			got, err := tt.reduce(event)
			if _, ok := err.(errors.InvalidArgument); !ok {
				t.Errorf("no wrong event mapping: %v, got: %v", err, got)
			}

			event = tt.args.event(t)
			got, err = tt.reduce(event)
			assertReduce(t, got, err, tt.want)
		})
	}
}