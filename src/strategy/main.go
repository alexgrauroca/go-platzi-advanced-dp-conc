package strategy

func Run() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("alex", "password", sha)
	passwordProtector.Hash()
	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}
