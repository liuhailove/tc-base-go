// @generated by protoc-gen-es v1.6.0 with parameter "target=ts"
// @generated from file tc_sip.proto (package tc, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tc.CreateSIPTrunkRequest
 */
export class CreateSIPTrunkRequest extends Message<CreateSIPTrunkRequest> {
  /**
   * CIDR or IPs that traffic is accepted from
   * An empty list means all inbound traffic is accepted.
   *
   * @generated from field: repeated string addresses = 1;
   */
  addresses: string[] = [];

  /**
   * `To` value that will be used when making a call
   *
   * @generated from field: string to = 2;
   */
  to = "";

  /**
   * Accepted `To` values. This Trunk will only accept a call made to
   * these numbers. This allows you to have distinct Trunks for different phone
   * numbers at the same provider.
   * An empty list means all dialed numbers are accepted.
   *
   * @generated from field: repeated string allowed_destinations_regex = 3;
   */
  allowedDestinationsRegex: string[] = [];

  constructor(data?: PartialMessage<CreateSIPTrunkRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.CreateSIPTrunkRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "addresses", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "allowed_destinations_regex", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPTrunkRequest {
    return new CreateSIPTrunkRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPTrunkRequest {
    return new CreateSIPTrunkRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPTrunkRequest {
    return new CreateSIPTrunkRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSIPTrunkRequest | PlainMessage<CreateSIPTrunkRequest> | undefined, b: CreateSIPTrunkRequest | PlainMessage<CreateSIPTrunkRequest> | undefined): boolean {
    return proto3.util.equals(CreateSIPTrunkRequest, a, b);
  }
}

/**
 * @generated from message tc.SIPTrunkInfo
 */
export class SIPTrunkInfo extends Message<SIPTrunkInfo> {
  /**
   * @generated from field: string sip_trunk_id = 1;
   */
  sipTrunkId = "";

  constructor(data?: PartialMessage<SIPTrunkInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPTrunkInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_trunk_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPTrunkInfo {
    return new SIPTrunkInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPTrunkInfo {
    return new SIPTrunkInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPTrunkInfo {
    return new SIPTrunkInfo().fromJsonString(jsonString, options);
  }

  static equals(a: SIPTrunkInfo | PlainMessage<SIPTrunkInfo> | undefined, b: SIPTrunkInfo | PlainMessage<SIPTrunkInfo> | undefined): boolean {
    return proto3.util.equals(SIPTrunkInfo, a, b);
  }
}

/**
 * @generated from message tc.ListSIPTrunkRequest
 */
export class ListSIPTrunkRequest extends Message<ListSIPTrunkRequest> {
  constructor(data?: PartialMessage<ListSIPTrunkRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.ListSIPTrunkRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPTrunkRequest {
    return new ListSIPTrunkRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPTrunkRequest {
    return new ListSIPTrunkRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPTrunkRequest {
    return new ListSIPTrunkRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListSIPTrunkRequest | PlainMessage<ListSIPTrunkRequest> | undefined, b: ListSIPTrunkRequest | PlainMessage<ListSIPTrunkRequest> | undefined): boolean {
    return proto3.util.equals(ListSIPTrunkRequest, a, b);
  }
}

/**
 * @generated from message tc.ListSIPTrunkResponse
 */
export class ListSIPTrunkResponse extends Message<ListSIPTrunkResponse> {
  /**
   * @generated from field: repeated tc.SIPTrunkInfo items = 1;
   */
  items: SIPTrunkInfo[] = [];

  constructor(data?: PartialMessage<ListSIPTrunkResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.ListSIPTrunkResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "items", kind: "message", T: SIPTrunkInfo, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPTrunkResponse {
    return new ListSIPTrunkResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPTrunkResponse {
    return new ListSIPTrunkResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPTrunkResponse {
    return new ListSIPTrunkResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListSIPTrunkResponse | PlainMessage<ListSIPTrunkResponse> | undefined, b: ListSIPTrunkResponse | PlainMessage<ListSIPTrunkResponse> | undefined): boolean {
    return proto3.util.equals(ListSIPTrunkResponse, a, b);
  }
}

/**
 * @generated from message tc.DeleteSIPTrunkRequest
 */
export class DeleteSIPTrunkRequest extends Message<DeleteSIPTrunkRequest> {
  /**
   * @generated from field: string sip_trunk_id = 1;
   */
  sipTrunkId = "";

  constructor(data?: PartialMessage<DeleteSIPTrunkRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.DeleteSIPTrunkRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_trunk_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSIPTrunkRequest {
    return new DeleteSIPTrunkRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSIPTrunkRequest {
    return new DeleteSIPTrunkRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSIPTrunkRequest {
    return new DeleteSIPTrunkRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSIPTrunkRequest | PlainMessage<DeleteSIPTrunkRequest> | undefined, b: DeleteSIPTrunkRequest | PlainMessage<DeleteSIPTrunkRequest> | undefined): boolean {
    return proto3.util.equals(DeleteSIPTrunkRequest, a, b);
  }
}

/**
 * @generated from message tc.SIPDispatchRuleDirect
 */
export class SIPDispatchRuleDirect extends Message<SIPDispatchRuleDirect> {
  /**
   * What room should call be directed into
   *
   * @generated from field: string room_name = 1;
   */
  roomName = "";

  /**
   * Optional pin required to enter room
   *
   * @generated from field: string pin = 2;
   */
  pin = "";

  constructor(data?: PartialMessage<SIPDispatchRuleDirect>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPDispatchRuleDirect";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "pin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleDirect {
    return new SIPDispatchRuleDirect().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleDirect {
    return new SIPDispatchRuleDirect().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleDirect {
    return new SIPDispatchRuleDirect().fromJsonString(jsonString, options);
  }

  static equals(a: SIPDispatchRuleDirect | PlainMessage<SIPDispatchRuleDirect> | undefined, b: SIPDispatchRuleDirect | PlainMessage<SIPDispatchRuleDirect> | undefined): boolean {
    return proto3.util.equals(SIPDispatchRuleDirect, a, b);
  }
}

/**
 * @generated from message tc.SIPDispatchRulePin
 */
export class SIPDispatchRulePin extends Message<SIPDispatchRulePin> {
  /**
   * What room should call be directed into
   *
   * @generated from field: string room_name = 1;
   */
  roomName = "";

  /**
   * Pin required to enter room
   *
   * @generated from field: string pin = 2;
   */
  pin = "";

  constructor(data?: PartialMessage<SIPDispatchRulePin>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPDispatchRulePin";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "pin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRulePin {
    return new SIPDispatchRulePin().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRulePin {
    return new SIPDispatchRulePin().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRulePin {
    return new SIPDispatchRulePin().fromJsonString(jsonString, options);
  }

  static equals(a: SIPDispatchRulePin | PlainMessage<SIPDispatchRulePin> | undefined, b: SIPDispatchRulePin | PlainMessage<SIPDispatchRulePin> | undefined): boolean {
    return proto3.util.equals(SIPDispatchRulePin, a, b);
  }
}

/**
 * @generated from message tc.SIPDispatchRuleIndividual
 */
export class SIPDispatchRuleIndividual extends Message<SIPDispatchRuleIndividual> {
  /**
   * Prefix used on new room name
   *
   * @generated from field: string room_prefix = 1;
   */
  roomPrefix = "";

  /**
   * Optional pin required to enter room
   *
   * @generated from field: string pin = 2;
   */
  pin = "";

  constructor(data?: PartialMessage<SIPDispatchRuleIndividual>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPDispatchRuleIndividual";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "room_prefix", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "pin", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleIndividual {
    return new SIPDispatchRuleIndividual().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleIndividual {
    return new SIPDispatchRuleIndividual().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleIndividual {
    return new SIPDispatchRuleIndividual().fromJsonString(jsonString, options);
  }

  static equals(a: SIPDispatchRuleIndividual | PlainMessage<SIPDispatchRuleIndividual> | undefined, b: SIPDispatchRuleIndividual | PlainMessage<SIPDispatchRuleIndividual> | undefined): boolean {
    return proto3.util.equals(SIPDispatchRuleIndividual, a, b);
  }
}

/**
 * @generated from message tc.CreateSIPDispatchRuleRequest
 */
export class CreateSIPDispatchRuleRequest extends Message<CreateSIPDispatchRuleRequest> {
  /**
   * @generated from oneof tc.CreateSIPDispatchRuleRequest.instrument
   */
  instrument: {
    /**
     * SIPDispatchRuleDirect is a `SIP Dispatch Rule` that puts a user directly into a room
     * This places users into an existing room. Optionally you can require a pin before a user can
     * enter the room
     *
     * @generated from field: tc.SIPDispatchRuleDirect dispatch_rule_direct = 1;
     */
    value: SIPDispatchRuleDirect;
    case: "dispatchRuleDirect";
  } | {
    /**
     * SIPDispatchRulePin is a `SIP Dispatch Rule` that allows a user to choose between multiple rooms.
     * The user is prompted for a pin and then can enter a individual room.
     *
     * @generated from field: tc.SIPDispatchRulePin dispatch_rule_pin = 2;
     */
    value: SIPDispatchRulePin;
    case: "dispatchRulePin";
  } | {
    /**
     * SIPDispatchRuleIndividual is a `SIP Dispatch Rule` that creates a new room for each caller.
     *
     * @generated from field: tc.SIPDispatchRuleIndividual dispatch_rule_individual = 3;
     */
    value: SIPDispatchRuleIndividual;
    case: "dispatchRuleIndividual";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * What trunks are accepted for this dispatch rule
   * If empty all trunks will match this dispatch rule
   *
   * @generated from field: repeated string trunk_ids = 4;
   */
  trunkIds: string[] = [];

  /**
   * By default the From value (Phone number) is used as the participant identity
   * If true a random value will be used instead
   *
   * @generated from field: bool hide_phone_number = 6;
   */
  hidePhoneNumber = false;

  constructor(data?: PartialMessage<CreateSIPDispatchRuleRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.CreateSIPDispatchRuleRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dispatch_rule_direct", kind: "message", T: SIPDispatchRuleDirect, oneof: "instrument" },
    { no: 2, name: "dispatch_rule_pin", kind: "message", T: SIPDispatchRulePin, oneof: "instrument" },
    { no: 3, name: "dispatch_rule_individual", kind: "message", T: SIPDispatchRuleIndividual, oneof: "instrument" },
    { no: 4, name: "trunk_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 6, name: "hide_phone_number", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPDispatchRuleRequest {
    return new CreateSIPDispatchRuleRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPDispatchRuleRequest {
    return new CreateSIPDispatchRuleRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPDispatchRuleRequest {
    return new CreateSIPDispatchRuleRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSIPDispatchRuleRequest | PlainMessage<CreateSIPDispatchRuleRequest> | undefined, b: CreateSIPDispatchRuleRequest | PlainMessage<CreateSIPDispatchRuleRequest> | undefined): boolean {
    return proto3.util.equals(CreateSIPDispatchRuleRequest, a, b);
  }
}

/**
 * @generated from message tc.SIPDispatchRuleInfo
 */
export class SIPDispatchRuleInfo extends Message<SIPDispatchRuleInfo> {
  /**
   * @generated from field: string sip_dispatch_rule_id = 1;
   */
  sipDispatchRuleId = "";

  constructor(data?: PartialMessage<SIPDispatchRuleInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPDispatchRuleInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_dispatch_rule_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPDispatchRuleInfo {
    return new SIPDispatchRuleInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPDispatchRuleInfo {
    return new SIPDispatchRuleInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPDispatchRuleInfo {
    return new SIPDispatchRuleInfo().fromJsonString(jsonString, options);
  }

  static equals(a: SIPDispatchRuleInfo | PlainMessage<SIPDispatchRuleInfo> | undefined, b: SIPDispatchRuleInfo | PlainMessage<SIPDispatchRuleInfo> | undefined): boolean {
    return proto3.util.equals(SIPDispatchRuleInfo, a, b);
  }
}

/**
 * @generated from message tc.ListSIPDispatchRuleRequest
 */
export class ListSIPDispatchRuleRequest extends Message<ListSIPDispatchRuleRequest> {
  constructor(data?: PartialMessage<ListSIPDispatchRuleRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.ListSIPDispatchRuleRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPDispatchRuleRequest {
    return new ListSIPDispatchRuleRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleRequest {
    return new ListSIPDispatchRuleRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleRequest {
    return new ListSIPDispatchRuleRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListSIPDispatchRuleRequest | PlainMessage<ListSIPDispatchRuleRequest> | undefined, b: ListSIPDispatchRuleRequest | PlainMessage<ListSIPDispatchRuleRequest> | undefined): boolean {
    return proto3.util.equals(ListSIPDispatchRuleRequest, a, b);
  }
}

/**
 * @generated from message tc.ListSIPDispatchRuleResponse
 */
export class ListSIPDispatchRuleResponse extends Message<ListSIPDispatchRuleResponse> {
  /**
   * @generated from field: repeated tc.SIPDispatchRuleInfo items = 1;
   */
  items: SIPDispatchRuleInfo[] = [];

  constructor(data?: PartialMessage<ListSIPDispatchRuleResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.ListSIPDispatchRuleResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "items", kind: "message", T: SIPDispatchRuleInfo, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPDispatchRuleResponse {
    return new ListSIPDispatchRuleResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleResponse {
    return new ListSIPDispatchRuleResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPDispatchRuleResponse {
    return new ListSIPDispatchRuleResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListSIPDispatchRuleResponse | PlainMessage<ListSIPDispatchRuleResponse> | undefined, b: ListSIPDispatchRuleResponse | PlainMessage<ListSIPDispatchRuleResponse> | undefined): boolean {
    return proto3.util.equals(ListSIPDispatchRuleResponse, a, b);
  }
}

/**
 * @generated from message tc.DeleteSIPDispatchRuleRequest
 */
export class DeleteSIPDispatchRuleRequest extends Message<DeleteSIPDispatchRuleRequest> {
  /**
   * @generated from field: string sip_dispatch_rule_id = 1;
   */
  sipDispatchRuleId = "";

  constructor(data?: PartialMessage<DeleteSIPDispatchRuleRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.DeleteSIPDispatchRuleRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_dispatch_rule_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSIPDispatchRuleRequest {
    return new DeleteSIPDispatchRuleRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSIPDispatchRuleRequest {
    return new DeleteSIPDispatchRuleRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSIPDispatchRuleRequest {
    return new DeleteSIPDispatchRuleRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSIPDispatchRuleRequest | PlainMessage<DeleteSIPDispatchRuleRequest> | undefined, b: DeleteSIPDispatchRuleRequest | PlainMessage<DeleteSIPDispatchRuleRequest> | undefined): boolean {
    return proto3.util.equals(DeleteSIPDispatchRuleRequest, a, b);
  }
}

/**
 * A SIP Participant is a singular SIP session connected to a LiveKit room via
 * a SIP Trunk into a SIP DispatchRule
 *
 * @generated from message tc.CreateSIPParticipantRequest
 */
export class CreateSIPParticipantRequest extends Message<CreateSIPParticipantRequest> {
  /**
   * What LiveKit room should this participant be connected too
   *
   * @generated from field: string room_name = 1;
   */
  roomName = "";

  /**
   * What SIP Trunk should be used to dial the user
   *
   * @generated from field: string sip_trunk_id = 2;
   */
  sipTrunkId = "";

  constructor(data?: PartialMessage<CreateSIPParticipantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.CreateSIPParticipantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "room_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "sip_trunk_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSIPParticipantRequest {
    return new CreateSIPParticipantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSIPParticipantRequest {
    return new CreateSIPParticipantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSIPParticipantRequest {
    return new CreateSIPParticipantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSIPParticipantRequest | PlainMessage<CreateSIPParticipantRequest> | undefined, b: CreateSIPParticipantRequest | PlainMessage<CreateSIPParticipantRequest> | undefined): boolean {
    return proto3.util.equals(CreateSIPParticipantRequest, a, b);
  }
}

/**
 * @generated from message tc.SIPParticipantInfo
 */
export class SIPParticipantInfo extends Message<SIPParticipantInfo> {
  /**
   * @generated from field: string sip_participant_id = 1;
   */
  sipParticipantId = "";

  constructor(data?: PartialMessage<SIPParticipantInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPParticipantInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_participant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPParticipantInfo {
    return new SIPParticipantInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPParticipantInfo {
    return new SIPParticipantInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPParticipantInfo {
    return new SIPParticipantInfo().fromJsonString(jsonString, options);
  }

  static equals(a: SIPParticipantInfo | PlainMessage<SIPParticipantInfo> | undefined, b: SIPParticipantInfo | PlainMessage<SIPParticipantInfo> | undefined): boolean {
    return proto3.util.equals(SIPParticipantInfo, a, b);
  }
}

/**
 * DTMF Request lets you send a DTMF message for a SIP Participant
 *
 * @generated from message tc.SendSIPParticipantDTMFRequest
 */
export class SendSIPParticipantDTMFRequest extends Message<SendSIPParticipantDTMFRequest> {
  /**
   * What SIP Participant to send this DTMF for
   *
   * @generated from field: string sip_participant_id = 1;
   */
  sipParticipantId = "";

  /**
   * Digits that will be sent via DTMF
   *
   * @generated from field: string digits = 2;
   */
  digits = "";

  constructor(data?: PartialMessage<SendSIPParticipantDTMFRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SendSIPParticipantDTMFRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_participant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "digits", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendSIPParticipantDTMFRequest {
    return new SendSIPParticipantDTMFRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendSIPParticipantDTMFRequest {
    return new SendSIPParticipantDTMFRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendSIPParticipantDTMFRequest {
    return new SendSIPParticipantDTMFRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SendSIPParticipantDTMFRequest | PlainMessage<SendSIPParticipantDTMFRequest> | undefined, b: SendSIPParticipantDTMFRequest | PlainMessage<SendSIPParticipantDTMFRequest> | undefined): boolean {
    return proto3.util.equals(SendSIPParticipantDTMFRequest, a, b);
  }
}

/**
 * @generated from message tc.SIPParticipantDTMFInfo
 */
export class SIPParticipantDTMFInfo extends Message<SIPParticipantDTMFInfo> {
  /**
   * @generated from field: string sip_participant_id = 1;
   */
  sipParticipantId = "";

  constructor(data?: PartialMessage<SIPParticipantDTMFInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.SIPParticipantDTMFInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_participant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SIPParticipantDTMFInfo {
    return new SIPParticipantDTMFInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SIPParticipantDTMFInfo {
    return new SIPParticipantDTMFInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SIPParticipantDTMFInfo {
    return new SIPParticipantDTMFInfo().fromJsonString(jsonString, options);
  }

  static equals(a: SIPParticipantDTMFInfo | PlainMessage<SIPParticipantDTMFInfo> | undefined, b: SIPParticipantDTMFInfo | PlainMessage<SIPParticipantDTMFInfo> | undefined): boolean {
    return proto3.util.equals(SIPParticipantDTMFInfo, a, b);
  }
}

/**
 * @generated from message tc.ListSIPParticipantRequest
 */
export class ListSIPParticipantRequest extends Message<ListSIPParticipantRequest> {
  constructor(data?: PartialMessage<ListSIPParticipantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.ListSIPParticipantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPParticipantRequest {
    return new ListSIPParticipantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPParticipantRequest {
    return new ListSIPParticipantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPParticipantRequest {
    return new ListSIPParticipantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListSIPParticipantRequest | PlainMessage<ListSIPParticipantRequest> | undefined, b: ListSIPParticipantRequest | PlainMessage<ListSIPParticipantRequest> | undefined): boolean {
    return proto3.util.equals(ListSIPParticipantRequest, a, b);
  }
}

/**
 * @generated from message tc.ListSIPParticipantResponse
 */
export class ListSIPParticipantResponse extends Message<ListSIPParticipantResponse> {
  /**
   * @generated from field: repeated tc.SIPParticipantInfo items = 1;
   */
  items: SIPParticipantInfo[] = [];

  constructor(data?: PartialMessage<ListSIPParticipantResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.ListSIPParticipantResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "items", kind: "message", T: SIPParticipantInfo, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListSIPParticipantResponse {
    return new ListSIPParticipantResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListSIPParticipantResponse {
    return new ListSIPParticipantResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListSIPParticipantResponse {
    return new ListSIPParticipantResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListSIPParticipantResponse | PlainMessage<ListSIPParticipantResponse> | undefined, b: ListSIPParticipantResponse | PlainMessage<ListSIPParticipantResponse> | undefined): boolean {
    return proto3.util.equals(ListSIPParticipantResponse, a, b);
  }
}

/**
 * @generated from message tc.DeleteSIPParticipantRequest
 */
export class DeleteSIPParticipantRequest extends Message<DeleteSIPParticipantRequest> {
  /**
   * @generated from field: string sip_participant_id = 1;
   */
  sipParticipantId = "";

  constructor(data?: PartialMessage<DeleteSIPParticipantRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "tc.DeleteSIPParticipantRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sip_participant_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteSIPParticipantRequest {
    return new DeleteSIPParticipantRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteSIPParticipantRequest {
    return new DeleteSIPParticipantRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteSIPParticipantRequest {
    return new DeleteSIPParticipantRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteSIPParticipantRequest | PlainMessage<DeleteSIPParticipantRequest> | undefined, b: DeleteSIPParticipantRequest | PlainMessage<DeleteSIPParticipantRequest> | undefined): boolean {
    return proto3.util.equals(DeleteSIPParticipantRequest, a, b);
  }
}

