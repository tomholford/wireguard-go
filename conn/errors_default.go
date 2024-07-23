//go:build !linux

/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2023 WireGuard LLC. All Rights Reserved.
 */

package conn

func errShouldDisableUDPGSO(_ error) bool {
	return false
}
