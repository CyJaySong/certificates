package provisioner

import (
	"github.com/smallstep/certificates/policy"
)

// newX509PolicyEngine creates a new x509 name policy engine
func newX509PolicyEngine(x509Opts *X509Options) (policy.X509NamePolicyEngine, error) {

	if x509Opts == nil {
		return nil, nil
	}

	options := []policy.NamePolicyOption{
		policy.WithSubjectCommonNameVerification(), // enable x509 Subject Common Name validation by default
	}

	allowed := x509Opts.GetAllowedNameOptions()
	if allowed != nil && allowed.HasNames() {
		options = append(options,
			policy.WithPermittedDNSDomains(allowed.DNSDomains),
			policy.WithPermittedIPsOrCIDRs(allowed.IPRanges),
			policy.WithPermittedEmailAddresses(allowed.EmailAddresses),
			policy.WithPermittedURIDomains(allowed.URIDomains),
		)
	}

	denied := x509Opts.GetDeniedNameOptions()
	if denied != nil && denied.HasNames() {
		options = append(options,
			policy.WithExcludedDNSDomains(denied.DNSDomains),
			policy.WithExcludedIPsOrCIDRs(denied.IPRanges),
			policy.WithExcludedEmailAddresses(denied.EmailAddresses),
			policy.WithExcludedURIDomains(denied.URIDomains),
		)
	}

	return policy.New(options...)
}

// newSSHPolicyEngine creates a new SSH name policy engine
func newSSHPolicyEngine(sshOpts *SSHOptions) (policy.SSHNamePolicyEngine, error) {

	if sshOpts == nil {
		return nil, nil
	}

	options := []policy.NamePolicyOption{}

	allowed := sshOpts.GetAllowedNameOptions()
	if allowed != nil && allowed.HasNames() {
		options = append(options,
			policy.WithPermittedDNSDomains(allowed.DNSDomains),
			policy.WithPermittedIPsOrCIDRs(allowed.IPRanges),
			policy.WithPermittedEmailAddresses(allowed.EmailAddresses),
			policy.WithPermittedPrincipals(allowed.Principals),
		)
	}

	denied := sshOpts.GetDeniedNameOptions()
	if denied != nil && denied.HasNames() {
		options = append(options,
			policy.WithExcludedDNSDomains(denied.DNSDomains),
			policy.WithExcludedIPsOrCIDRs(denied.IPRanges),
			policy.WithExcludedEmailAddresses(denied.EmailAddresses),
			policy.WithExcludedPrincipals(denied.Principals),
		)
	}

	return policy.New(options...)
}