package server

import (
	"github.com/meerkat-chain/mchain/consensus"
	consensusDev "github.com/meerkat-chain/mchain/consensus/dev"
	consensusDummy "github.com/meerkat-chain/mchain/consensus/dummy"
	consensusIBFT "github.com/meerkat-chain/mchain/consensus/ibft"
	"github.com/meerkat-chain/mchain/secrets"
	"github.com/meerkat-chain/mchain/secrets/awsssm"
	"github.com/meerkat-chain/mchain/secrets/gcpssm"
	"github.com/meerkat-chain/mchain/secrets/hashicorpvault"
	"github.com/meerkat-chain/mchain/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
