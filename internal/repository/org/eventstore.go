package org

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(OrgAddedEventType, OrgAddedEventMapper).
		RegisterFilterEventMapper(OrgChangedEventType, OrgChangedEventMapper).
		RegisterFilterEventMapper(OrgDeactivatedEventType, OrgDeactivatedEventMapper).
		RegisterFilterEventMapper(OrgReactivatedEventType, OrgReactivatedEventMapper).
		RegisterFilterEventMapper(OrgDomainAddedEventType, DomainAddedEventMapper).
		RegisterFilterEventMapper(OrgDomainVerificationAddedEventType, DomainVerificationAddedEventMapper).
		RegisterFilterEventMapper(OrgDomainVerificationFailedEventType, DomainVerificationFailedEventMapper).
		RegisterFilterEventMapper(OrgDomainVerifiedEventType, DomainVerifiedEventMapper).
		RegisterFilterEventMapper(OrgDomainPrimarySetEventType, DomainPrimarySetEventMapper).
		RegisterFilterEventMapper(OrgDomainRemovedEventType, DomainRemovedEventMapper).
		RegisterFilterEventMapper(MemberAddedEventType, MemberAddedEventMapper).
		RegisterFilterEventMapper(MemberChangedEventType, MemberChangedEventMapper).
		RegisterFilterEventMapper(MemberRemovedEventType, MemberRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyAddedEventType, LabelPolicyAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyChangedEventType, LabelPolicyChangedEventMapper).
		RegisterFilterEventMapper(LabelPolicyRemovedEventType, LabelPolicyRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyAddedEventType, LoginPolicyAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyChangedEventType, LoginPolicyChangedEventMapper).
		RegisterFilterEventMapper(LoginPolicyRemovedEventType, LoginPolicyRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorAddedEventType, SecondFactorAddedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorRemovedEventType, SecondFactorRemovedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorAddedEventType, MultiFactorAddedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorRemovedEventType, MultiFactorRemovedEventEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderAddedEventType, IdentityProviderAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderRemovedEventType, IdentityProviderRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderCascadeRemovedEventType, IdentityProviderCascadeRemovedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyAddedEventType, OrgIAMPolicyAddedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyChangedEventType, OrgIAMPolicyChangedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyRemovedEventType, OrgIAMPolicyRemovedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyAddedEventType, PasswordAgePolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyChangedEventType, PasswordAgePolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyRemovedEventType, PasswordAgePolicyRemovedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyAddedEventType, PasswordComplexityPolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyChangedEventType, PasswordComplexityPolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyRemovedEventType, PasswordComplexityPolicyRemovedEventMapper).
		RegisterFilterEventMapper(PasswordLockoutPolicyAddedEventType, PasswordLockoutPolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordLockoutPolicyChangedEventType, PasswordLockoutPolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordLockoutPolicyRemovedEventType, PasswordLockoutPolicyRemovedEventMapper).
		RegisterFilterEventMapper(MailTemplateAddedEventType, MailTemplateAddedEventMapper).
		RegisterFilterEventMapper(MailTemplateChangedEventType, MailTemplateChangedEventMapper).
		RegisterFilterEventMapper(MailTemplateRemovedEventType, MailTemplateRemovedEventMapper).
		RegisterFilterEventMapper(MailTextAddedEventType, MailTextAddedEventMapper).
		RegisterFilterEventMapper(MailTextChangedEventType, MailTextChangedEventMapper).
		RegisterFilterEventMapper(MailTextRemovedEventType, MailTextRemovedEventMapper).
		RegisterFilterEventMapper(CustomTextSetEventType, CustomTextSetEventMapper).
		RegisterFilterEventMapper(CustomTextRemovedEventType, CustomTextRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigAddedEventType, IDPConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPConfigChangedEventType, IDPConfigChangedEventMapper).
		RegisterFilterEventMapper(IDPConfigRemovedEventType, IDPConfigRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigDeactivatedEventType, IDPConfigDeactivatedEventMapper).
		RegisterFilterEventMapper(IDPConfigReactivatedEventType, IDPConfigReactivatedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigAddedEventType, IDPOIDCConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigChangedEventType, IDPOIDCConfigChangedEventMapper).
		RegisterFilterEventMapper(FeaturesSetEventType, FeaturesSetEventMapper).
		RegisterFilterEventMapper(FeaturesRemovedEventType, FeaturesRemovedEventMapper)
}
