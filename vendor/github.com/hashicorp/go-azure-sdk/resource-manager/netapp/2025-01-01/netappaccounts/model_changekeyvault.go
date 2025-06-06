package netappaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeKeyVault struct {
	KeyName                  string                    `json:"keyName"`
	KeyVaultPrivateEndpoints []KeyVaultPrivateEndpoint `json:"keyVaultPrivateEndpoints"`
	KeyVaultResourceId       *string                   `json:"keyVaultResourceId,omitempty"`
	KeyVaultUri              string                    `json:"keyVaultUri"`
}
