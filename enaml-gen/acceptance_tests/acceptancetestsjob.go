package acceptance_tests 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type AcceptanceTestsJob struct {

	/*AcceptanceTests - Descr: URL of the Cloud Controller API Default: <nil>
*/
	AcceptanceTests *AcceptanceTests `yaml:"acceptance_tests,omitempty"`

	/*TcpEmitter - Descr: Password for UAA client for the tcp emitter. Default: <nil>
*/
	TcpEmitter *TcpEmitter `yaml:"tcp_emitter,omitempty"`

}