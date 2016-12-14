package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Jwt struct {

	/*Policy - Descr: The global refresh token validity for all zones if nothing is configured on the client Default: 2592000
*/
	Policy *JwtPolicy `yaml:"policy,omitempty"`

	/*Refresh - Descr: Disallows refresh-token grant for any client for which the user has not approved the `uaa.offline_token` scope Default: false
*/
	Refresh *Refresh `yaml:"refresh,omitempty"`

	/*Revocable - Descr: Set to true if you wish that even JWT tokens become individually revocable and stored in the UAA token storage. This setting applies to the default zone only. Default: false
*/
	Revocable interface{} `yaml:"revocable,omitempty"`

	/*VerificationKey - Descr: Deprecated. Use uaa.jwt.policy.keys. The key used to verify JWT-based OAuth2 tokens Default: <nil>
*/
	VerificationKey interface{} `yaml:"verification_key,omitempty"`

	/*Claims - Descr: List of claims to exclude from the JWT-based OAuth2 tokens Default: <nil>
*/
	Claims *Claims `yaml:"claims,omitempty"`

	/*SigningKey - Descr: Deprecated. Use uaa.jwt.policy.keys. The key used to sign the JWT-based OAuth2 tokens Default: <nil>
*/
	SigningKey interface{} `yaml:"signing_key,omitempty"`

}