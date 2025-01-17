// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;
use wasmlib::host::*;

use crate::*;
use crate::keys::*;

#[derive(Clone, Copy)]
pub struct ImmutableGetFactorResults {
    pub(crate) id: i32,
}

impl ImmutableGetFactorResults {
    pub fn factor(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_FACTOR))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetFactorResults {
    pub(crate) id: i32,
}

impl MutableGetFactorResults {
    pub fn factor(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_FACTOR))
    }
}

#[derive(Clone, Copy)]
pub struct ImmutableGetOwnerResults {
    pub(crate) id: i32,
}

impl ImmutableGetOwnerResults {
    pub fn owner(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_RESULT_OWNER))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetOwnerResults {
    pub(crate) id: i32,
}

impl MutableGetOwnerResults {
    pub fn owner(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_RESULT_OWNER))
    }
}
