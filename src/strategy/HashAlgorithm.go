package strategy

type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}
