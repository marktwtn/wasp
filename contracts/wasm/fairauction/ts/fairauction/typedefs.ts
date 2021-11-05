// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib"
import * as sc from "./index";

export class ArrayOfImmutableAgentID {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getAgentID(index: i32): wasmlib.ScImmutableAgentID {
        return new wasmlib.ScImmutableAgentID(this.objID, new wasmlib.Key32(index));
    }
}

export class ImmutableBidderList extends ArrayOfImmutableAgentID {
};

export class ArrayOfMutableAgentID {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID);
    }

    length(): i32 {
        return wasmlib.getLength(this.objID);
    }

    getAgentID(index: i32): wasmlib.ScMutableAgentID {
        return new wasmlib.ScMutableAgentID(this.objID, new wasmlib.Key32(index));
    }
}

export class MutableBidderList extends ArrayOfMutableAgentID {
};

export class MapAgentIDToImmutableBid {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    getBid(key: wasmlib.ScAgentID): sc.ImmutableBid {
        return new sc.ImmutableBid(this.objID, key.getKeyID());
    }
}

export class ImmutableBids extends MapAgentIDToImmutableBid {
};

export class MapAgentIDToMutableBid {
    objID: i32;

    constructor(objID: i32) {
        this.objID = objID;
    }

    clear(): void {
        wasmlib.clear(this.objID)
    }

    getBid(key: wasmlib.ScAgentID): sc.MutableBid {
        return new sc.MutableBid(this.objID, key.getKeyID());
    }
}

export class MutableBids extends MapAgentIDToMutableBid {
};