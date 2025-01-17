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
use crate::structs::*;

#[derive(Clone, Copy)]
pub struct ImmutableGetInfoResults {
    pub(crate) id: i32,
}

impl ImmutableGetInfoResults {
    pub fn bidders(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, idx_map(IDX_RESULT_BIDDERS))
    }

    pub fn color(&self) -> ScImmutableColor {
        ScImmutableColor::new(self.id, idx_map(IDX_RESULT_COLOR))
    }

    pub fn creator(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_RESULT_CREATOR))
    }

    pub fn deposit(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_DEPOSIT))
    }

    pub fn description(&self) -> ScImmutableString {
        ScImmutableString::new(self.id, idx_map(IDX_RESULT_DESCRIPTION))
    }

    pub fn duration(&self) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.id, idx_map(IDX_RESULT_DURATION))
    }

    pub fn highest_bid(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_HIGHEST_BID))
    }

    pub fn highest_bidder(&self) -> ScImmutableAgentID {
        ScImmutableAgentID::new(self.id, idx_map(IDX_RESULT_HIGHEST_BIDDER))
    }

    pub fn minimum_bid(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_MINIMUM_BID))
    }

    pub fn num_tokens(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_NUM_TOKENS))
    }

    pub fn owner_margin(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_OWNER_MARGIN))
    }

    pub fn when_started(&self) -> ScImmutableInt64 {
        ScImmutableInt64::new(self.id, idx_map(IDX_RESULT_WHEN_STARTED))
    }
}

#[derive(Clone, Copy)]
pub struct MutableGetInfoResults {
    pub(crate) id: i32,
}

impl MutableGetInfoResults {
    pub fn bidders(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, idx_map(IDX_RESULT_BIDDERS))
    }

    pub fn color(&self) -> ScMutableColor {
        ScMutableColor::new(self.id, idx_map(IDX_RESULT_COLOR))
    }

    pub fn creator(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_RESULT_CREATOR))
    }

    pub fn deposit(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_DEPOSIT))
    }

    pub fn description(&self) -> ScMutableString {
        ScMutableString::new(self.id, idx_map(IDX_RESULT_DESCRIPTION))
    }

    pub fn duration(&self) -> ScMutableInt32 {
        ScMutableInt32::new(self.id, idx_map(IDX_RESULT_DURATION))
    }

    pub fn highest_bid(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_HIGHEST_BID))
    }

    pub fn highest_bidder(&self) -> ScMutableAgentID {
        ScMutableAgentID::new(self.id, idx_map(IDX_RESULT_HIGHEST_BIDDER))
    }

    pub fn minimum_bid(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_MINIMUM_BID))
    }

    pub fn num_tokens(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_NUM_TOKENS))
    }

    pub fn owner_margin(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_OWNER_MARGIN))
    }

    pub fn when_started(&self) -> ScMutableInt64 {
        ScMutableInt64::new(self.id, idx_map(IDX_RESULT_WHEN_STARTED))
    }
}
