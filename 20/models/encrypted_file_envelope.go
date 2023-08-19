package models

import c "aoc/common"

type encrypted_file_envelope struct {
	encrypted_file EncryptedFile
}

func EncryptedFileEnvelope(encrypted_file EncryptedFile) c.Envelope[EncryptedFile] {
	return encrypted_file_envelope{encrypted_file: encrypted_file}
}

func (envelope encrypted_file_envelope) Get() EncryptedFile {
	return c.ShallowCopy(envelope.encrypted_file)
}

func EncryptedFileEnvelopeEqualityFunction(lhs, rhs c.Envelope[EncryptedFile]) bool {
	return c.ArrayEqual(lhs.Get(), rhs.Get())
}
