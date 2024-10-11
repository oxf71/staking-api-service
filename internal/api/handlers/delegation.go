package handlers

import (
	"net/http"

	"github.com/babylonlabs-io/staking-api-service/internal/services"
	"github.com/babylonlabs-io/staking-api-service/internal/types"
)

// GetDelegationByTxHash @Summary Get a delegation
// @Description Retrieves a delegation by a given transaction hash
// @Produce json
// @Param staking_tx_hash_hex query string true "Staking transaction hash in hex format"
// @Success 200 {object} PublicResponse[services.DelegationPublic] "Delegation"
// @Failure 400 {object} types.Error "Error: Bad Request"
// @Router /v1/delegation [get]
func (h *Handler) GetDelegationByTxHash(request *http.Request) (*Result, *types.Error) {
	stakingTxHash, err := parseTxHashQuery(request, "staking_tx_hash_hex")
	if err != nil {
		return nil, err
	}
	delegation, err := h.services.GetDelegation(request.Context(), stakingTxHash)
	if err != nil {
		return nil, err
	}

	return NewResult(services.FromDelegationDocument(delegation)), nil
}

func (h *Handler) GetDelegationByFP(request *http.Request) (*Result, *types.Error) {
	fpPk, err := parsePublicKeyQuery(request, "fp_btc_pk", true)
	if err != nil {
		return nil, err
	}
	paginationKey, err := parsePaginationQuery(request)
	if err != nil {
		return nil, err
	}
	stateFilter, err := parseStateFilterQuery(request, "state")
	if err != nil {
		return nil, err
	}
	// TODO: height filter
	delegations, newPaginationKey, err := h.services.DelegationsByFP(
		request.Context(), fpPk, stateFilter, paginationKey,
	)
	if err != nil {
		return nil, err
	}

	return NewResultWithPagination(delegations, newPaginationKey), nil
}
