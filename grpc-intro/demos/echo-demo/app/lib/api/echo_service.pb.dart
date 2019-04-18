///
//  Generated code. Do not modify.
//  source: echo_service.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, Map, override;

import 'package:protobuf/protobuf.dart' as $pb;

class EchoMessage extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('EchoMessage', package: const $pb.PackageName('echo'))
    ..aOS(1, 'value')
    ..aOB(2, 'reverse')
    ..hasRequiredFields = false
  ;

  EchoMessage() : super();
  EchoMessage.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  EchoMessage.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  EchoMessage clone() => new EchoMessage()..mergeFromMessage(this);
  EchoMessage copyWith(void Function(EchoMessage) updates) => super.copyWith((message) => updates(message as EchoMessage));
  $pb.BuilderInfo get info_ => _i;
  static EchoMessage create() => new EchoMessage();
  EchoMessage createEmptyInstance() => create();
  static $pb.PbList<EchoMessage> createRepeated() => new $pb.PbList<EchoMessage>();
  static EchoMessage getDefault() => _defaultInstance ??= create()..freeze();
  static EchoMessage _defaultInstance;
  static void $checkItem(EchoMessage v) {
    if (v is! EchoMessage) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  String get value => $_getS(0, '');
  set value(String v) { $_setString(0, v); }
  bool hasValue() => $_has(0);
  void clearValue() => clearField(1);

  bool get reverse => $_get(1, false);
  set reverse(bool v) { $_setBool(1, v); }
  bool hasReverse() => $_has(1);
  void clearReverse() => clearField(2);
}

