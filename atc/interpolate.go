// Code generated by github.com/concourse/concourse/vars/interp/interpgen. Do not modify!

package atc

import (
	"encoding/json"

	"github.com/concourse/concourse/vars/interp"
)

type interpCPULimit struct {
	I interface {
		Interpolate(interp.Resolver) (CPULimit, error)
	}
}
type interpCPULimitVar interp.Var

func (v CPULimit) Interpolate(interp.Resolver) (CPULimit, error) {
	return CPULimit(v), nil
}
func (v CPULimit) Wrap() interpCPULimit     { return interpCPULimit{I: v} }
func (v CPULimit) WrapPtr() *interpCPULimit { return &interpCPULimit{I: v} }
func (v interpCPULimitVar) Interpolate(r interp.Resolver) (CPULimit, error) {
	var dst CPULimit
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpCPULimit) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpCPULimit) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpCPULimitVar(v)
		return nil
	}
	var raw CPULimit
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = CPULimit(raw)
	return nil
}

type interpDuration struct {
	I interface {
		Interpolate(interp.Resolver) (Duration, error)
	}
}
type interpDurationVar interp.Var

func (v Duration) Interpolate(interp.Resolver) (Duration, error) {
	return Duration(v), nil
}
func (v Duration) Wrap() interpDuration     { return interpDuration{I: v} }
func (v Duration) WrapPtr() *interpDuration { return &interpDuration{I: v} }
func (v interpDurationVar) Interpolate(r interp.Resolver) (Duration, error) {
	var dst Duration
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpDuration) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpDuration) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpDurationVar(v)
		return nil
	}
	var raw Duration
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = Duration(raw)
	return nil
}

type interpInputsConfig struct {
	I interface {
		Interpolate(interp.Resolver) (InputsConfig, error)
	}
}
type interpInputsConfigVar interp.Var

func (v InputsConfig) Interpolate(interp.Resolver) (InputsConfig, error) {
	return InputsConfig(v), nil
}
func (v InputsConfig) Wrap() interpInputsConfig     { return interpInputsConfig{I: v} }
func (v InputsConfig) WrapPtr() *interpInputsConfig { return &interpInputsConfig{I: v} }
func (v interpInputsConfigVar) Interpolate(r interp.Resolver) (InputsConfig, error) {
	var dst InputsConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInputsConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInputsConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInputsConfigVar(v)
		return nil
	}
	var raw InputsConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InputsConfig(raw)
	return nil
}

type interpInterpArtifactMapping struct {
	I interface {
		Interpolate(interp.Resolver) (InterpArtifactMapping, error)
	}
}
type interpInterpArtifactMappingVar interp.Var

func (v InterpArtifactMapping) Interpolate(interp.Resolver) (InterpArtifactMapping, error) {
	return InterpArtifactMapping(v), nil
}
func (v InterpArtifactMapping) Wrap() interpInterpArtifactMapping {
	return interpInterpArtifactMapping{I: v}
}
func (v InterpArtifactMapping) WrapPtr() *interpInterpArtifactMapping {
	return &interpInterpArtifactMapping{I: v}
}
func (v interpInterpArtifactMappingVar) Interpolate(r interp.Resolver) (InterpArtifactMapping, error) {
	var dst InterpArtifactMapping
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpArtifactMapping) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpArtifactMapping) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpArtifactMappingVar(v)
		return nil
	}
	var raw InterpArtifactMapping
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpArtifactMapping(raw)
	return nil
}

type interpInterpContainerLimits struct {
	I interface {
		Interpolate(interp.Resolver) (interpContainerLimits, error)
	}
}
type interpInterpContainerLimitsVar interp.Var

func (v interpContainerLimits) Interpolate(interp.Resolver) (interpContainerLimits, error) {
	return interpContainerLimits(v), nil
}
func (v interpContainerLimits) Wrap() interpInterpContainerLimits {
	return interpInterpContainerLimits{I: v}
}
func (v interpContainerLimits) WrapPtr() *interpInterpContainerLimits {
	return &interpInterpContainerLimits{I: v}
}
func (v interpInterpContainerLimitsVar) Interpolate(r interp.Resolver) (interpContainerLimits, error) {
	var dst interpContainerLimits
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpContainerLimits) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpContainerLimits) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpContainerLimitsVar(v)
		return nil
	}
	var raw interpContainerLimits
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpContainerLimits(raw)
	return nil
}

type interpInterpImageResource struct {
	I interface {
		Interpolate(interp.Resolver) (InterpImageResource, error)
	}
}
type interpInterpImageResourceVar interp.Var

func (v InterpImageResource) Interpolate(interp.Resolver) (InterpImageResource, error) {
	return InterpImageResource(v), nil
}
func (v InterpImageResource) Wrap() interpInterpImageResource {
	return interpInterpImageResource{I: v}
}
func (v InterpImageResource) WrapPtr() *interpInterpImageResource {
	return &interpInterpImageResource{I: v}
}
func (v interpInterpImageResourceVar) Interpolate(r interp.Resolver) (InterpImageResource, error) {
	var dst InterpImageResource
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpImageResource) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpImageResource) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpImageResourceVar(v)
		return nil
	}
	var raw InterpImageResource
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpImageResource(raw)
	return nil
}

type interpInterpParams struct {
	I interface {
		Interpolate(interp.Resolver) (InterpParams, error)
	}
}
type interpInterpParamsVar interp.Var

func (v InterpParams) Interpolate(interp.Resolver) (InterpParams, error) {
	return InterpParams(v), nil
}
func (v InterpParams) Wrap() interpInterpParams     { return interpInterpParams{I: v} }
func (v InterpParams) WrapPtr() *interpInterpParams { return &interpInterpParams{I: v} }
func (v interpInterpParamsVar) Interpolate(r interp.Resolver) (InterpParams, error) {
	var dst InterpParams
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpParams) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpParams) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpParamsVar(v)
		return nil
	}
	var raw InterpParams
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpParams(raw)
	return nil
}

type interpInterpSource struct {
	I interface {
		Interpolate(interp.Resolver) (InterpSource, error)
	}
}
type interpInterpSourceVar interp.Var

func (v InterpSource) Interpolate(interp.Resolver) (InterpSource, error) {
	return InterpSource(v), nil
}
func (v InterpSource) Wrap() interpInterpSource     { return interpInterpSource{I: v} }
func (v InterpSource) WrapPtr() *interpInterpSource { return &interpInterpSource{I: v} }
func (v interpInterpSourceVar) Interpolate(r interp.Resolver) (InterpSource, error) {
	var dst InterpSource
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpSource) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpSource) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpSourceVar(v)
		return nil
	}
	var raw InterpSource
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpSource(raw)
	return nil
}

type interpInterpStringList struct {
	I interface {
		Interpolate(interp.Resolver) (InterpStringList, error)
	}
}
type interpInterpStringListVar interp.Var

func (v InterpStringList) Interpolate(interp.Resolver) (InterpStringList, error) {
	return InterpStringList(v), nil
}
func (v InterpStringList) Wrap() interpInterpStringList {
	return interpInterpStringList{I: v}
}
func (v InterpStringList) WrapPtr() *interpInterpStringList {
	return &interpInterpStringList{I: v}
}
func (v interpInterpStringListVar) Interpolate(r interp.Resolver) (InterpStringList, error) {
	var dst InterpStringList
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpStringList) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpStringList) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpStringListVar(v)
		return nil
	}
	var raw InterpStringList
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpStringList(raw)
	return nil
}

type interpInterpTags struct {
	I interface {
		Interpolate(interp.Resolver) (InterpTags, error)
	}
}
type interpInterpTagsVar interp.Var

func (v InterpTags) Interpolate(interp.Resolver) (InterpTags, error) {
	return InterpTags(v), nil
}
func (v InterpTags) Wrap() interpInterpTags     { return interpInterpTags{I: v} }
func (v InterpTags) WrapPtr() *interpInterpTags { return &interpInterpTags{I: v} }
func (v interpInterpTagsVar) Interpolate(r interp.Resolver) (InterpTags, error) {
	var dst InterpTags
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTags) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTags) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTagsVar(v)
		return nil
	}
	var raw InterpTags
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpTags(raw)
	return nil
}

type interpInterpTaskCacheConfig struct {
	I interface {
		Interpolate(interp.Resolver) (interpTaskCacheConfig, error)
	}
}
type interpInterpTaskCacheConfigVar interp.Var

func (v interpTaskCacheConfig) Interpolate(interp.Resolver) (interpTaskCacheConfig, error) {
	return interpTaskCacheConfig(v), nil
}
func (v interpTaskCacheConfig) Wrap() interpInterpTaskCacheConfig {
	return interpInterpTaskCacheConfig{I: v}
}
func (v interpTaskCacheConfig) WrapPtr() *interpInterpTaskCacheConfig {
	return &interpInterpTaskCacheConfig{I: v}
}
func (v interpInterpTaskCacheConfigVar) Interpolate(r interp.Resolver) (interpTaskCacheConfig, error) {
	var dst interpTaskCacheConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskCacheConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskCacheConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskCacheConfigVar(v)
		return nil
	}
	var raw interpTaskCacheConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpTaskCacheConfig(raw)
	return nil
}

type interpInterpTaskCacheConfigs struct {
	I interface {
		Interpolate(interp.Resolver) (interpTaskCacheConfigs, error)
	}
}
type interpInterpTaskCacheConfigsVar interp.Var

func (v interpTaskCacheConfigs) Interpolate(interp.Resolver) (interpTaskCacheConfigs, error) {
	return interpTaskCacheConfigs(v), nil
}
func (v interpTaskCacheConfigs) Wrap() interpInterpTaskCacheConfigs {
	return interpInterpTaskCacheConfigs{I: v}
}
func (v interpTaskCacheConfigs) WrapPtr() *interpInterpTaskCacheConfigs {
	return &interpInterpTaskCacheConfigs{I: v}
}
func (v interpInterpTaskCacheConfigsVar) Interpolate(r interp.Resolver) (interpTaskCacheConfigs, error) {
	var dst interpTaskCacheConfigs
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskCacheConfigs) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskCacheConfigs) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskCacheConfigsVar(v)
		return nil
	}
	var raw interpTaskCacheConfigs
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpTaskCacheConfigs(raw)
	return nil
}

type interpInterpTaskConfig struct {
	I interface {
		Interpolate(interp.Resolver) (InterpTaskConfig, error)
	}
}
type interpInterpTaskConfigVar interp.Var

func (v InterpTaskConfig) Interpolate(interp.Resolver) (InterpTaskConfig, error) {
	return InterpTaskConfig(v), nil
}
func (v InterpTaskConfig) Wrap() interpInterpTaskConfig {
	return interpInterpTaskConfig{I: v}
}
func (v InterpTaskConfig) WrapPtr() *interpInterpTaskConfig {
	return &interpInterpTaskConfig{I: v}
}
func (v interpInterpTaskConfigVar) Interpolate(r interp.Resolver) (InterpTaskConfig, error) {
	var dst InterpTaskConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskConfigVar(v)
		return nil
	}
	var raw InterpTaskConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpTaskConfig(raw)
	return nil
}

type interpInterpTaskEnv struct {
	I interface {
		Interpolate(interp.Resolver) (InterpTaskEnv, error)
	}
}
type interpInterpTaskEnvVar interp.Var

func (v InterpTaskEnv) Interpolate(interp.Resolver) (InterpTaskEnv, error) {
	return InterpTaskEnv(v), nil
}
func (v InterpTaskEnv) Wrap() interpInterpTaskEnv     { return interpInterpTaskEnv{I: v} }
func (v InterpTaskEnv) WrapPtr() *interpInterpTaskEnv { return &interpInterpTaskEnv{I: v} }
func (v interpInterpTaskEnvVar) Interpolate(r interp.Resolver) (InterpTaskEnv, error) {
	var dst InterpTaskEnv
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskEnv) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskEnv) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskEnvVar(v)
		return nil
	}
	var raw InterpTaskEnv
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpTaskEnv(raw)
	return nil
}

type interpInterpTaskInputConfig struct {
	I interface {
		Interpolate(interp.Resolver) (interpTaskInputConfig, error)
	}
}
type interpInterpTaskInputConfigVar interp.Var

func (v interpTaskInputConfig) Interpolate(interp.Resolver) (interpTaskInputConfig, error) {
	return interpTaskInputConfig(v), nil
}
func (v interpTaskInputConfig) Wrap() interpInterpTaskInputConfig {
	return interpInterpTaskInputConfig{I: v}
}
func (v interpTaskInputConfig) WrapPtr() *interpInterpTaskInputConfig {
	return &interpInterpTaskInputConfig{I: v}
}
func (v interpInterpTaskInputConfigVar) Interpolate(r interp.Resolver) (interpTaskInputConfig, error) {
	var dst interpTaskInputConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskInputConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskInputConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskInputConfigVar(v)
		return nil
	}
	var raw interpTaskInputConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpTaskInputConfig(raw)
	return nil
}

type interpInterpTaskInputConfigs struct {
	I interface {
		Interpolate(interp.Resolver) (interpTaskInputConfigs, error)
	}
}
type interpInterpTaskInputConfigsVar interp.Var

func (v interpTaskInputConfigs) Interpolate(interp.Resolver) (interpTaskInputConfigs, error) {
	return interpTaskInputConfigs(v), nil
}
func (v interpTaskInputConfigs) Wrap() interpInterpTaskInputConfigs {
	return interpInterpTaskInputConfigs{I: v}
}
func (v interpTaskInputConfigs) WrapPtr() *interpInterpTaskInputConfigs {
	return &interpInterpTaskInputConfigs{I: v}
}
func (v interpInterpTaskInputConfigsVar) Interpolate(r interp.Resolver) (interpTaskInputConfigs, error) {
	var dst interpTaskInputConfigs
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskInputConfigs) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskInputConfigs) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskInputConfigsVar(v)
		return nil
	}
	var raw interpTaskInputConfigs
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpTaskInputConfigs(raw)
	return nil
}

type interpInterpTaskOutputConfig struct {
	I interface {
		Interpolate(interp.Resolver) (interpTaskOutputConfig, error)
	}
}
type interpInterpTaskOutputConfigVar interp.Var

func (v interpTaskOutputConfig) Interpolate(interp.Resolver) (interpTaskOutputConfig, error) {
	return interpTaskOutputConfig(v), nil
}
func (v interpTaskOutputConfig) Wrap() interpInterpTaskOutputConfig {
	return interpInterpTaskOutputConfig{I: v}
}
func (v interpTaskOutputConfig) WrapPtr() *interpInterpTaskOutputConfig {
	return &interpInterpTaskOutputConfig{I: v}
}
func (v interpInterpTaskOutputConfigVar) Interpolate(r interp.Resolver) (interpTaskOutputConfig, error) {
	var dst interpTaskOutputConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskOutputConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskOutputConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskOutputConfigVar(v)
		return nil
	}
	var raw interpTaskOutputConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpTaskOutputConfig(raw)
	return nil
}

type interpInterpTaskOutputConfigs struct {
	I interface {
		Interpolate(interp.Resolver) (interpTaskOutputConfigs, error)
	}
}
type interpInterpTaskOutputConfigsVar interp.Var

func (v interpTaskOutputConfigs) Interpolate(interp.Resolver) (interpTaskOutputConfigs, error) {
	return interpTaskOutputConfigs(v), nil
}
func (v interpTaskOutputConfigs) Wrap() interpInterpTaskOutputConfigs {
	return interpInterpTaskOutputConfigs{I: v}
}
func (v interpTaskOutputConfigs) WrapPtr() *interpInterpTaskOutputConfigs {
	return &interpInterpTaskOutputConfigs{I: v}
}
func (v interpInterpTaskOutputConfigsVar) Interpolate(r interp.Resolver) (interpTaskOutputConfigs, error) {
	var dst interpTaskOutputConfigs
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskOutputConfigs) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskOutputConfigs) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskOutputConfigsVar(v)
		return nil
	}
	var raw interpTaskOutputConfigs
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = interpTaskOutputConfigs(raw)
	return nil
}

type interpInterpTaskRunConfig struct {
	I interface {
		Interpolate(interp.Resolver) (InterpTaskRunConfig, error)
	}
}
type interpInterpTaskRunConfigVar interp.Var

func (v InterpTaskRunConfig) Interpolate(interp.Resolver) (InterpTaskRunConfig, error) {
	return InterpTaskRunConfig(v), nil
}
func (v InterpTaskRunConfig) Wrap() interpInterpTaskRunConfig {
	return interpInterpTaskRunConfig{I: v}
}
func (v InterpTaskRunConfig) WrapPtr() *interpInterpTaskRunConfig {
	return &interpInterpTaskRunConfig{I: v}
}
func (v interpInterpTaskRunConfigVar) Interpolate(r interp.Resolver) (InterpTaskRunConfig, error) {
	var dst InterpTaskRunConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpInterpTaskRunConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpInterpTaskRunConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpInterpTaskRunConfigVar(v)
		return nil
	}
	var raw InterpTaskRunConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = InterpTaskRunConfig(raw)
	return nil
}

type interpMaxInFlightConfig struct {
	I interface {
		Interpolate(interp.Resolver) (MaxInFlightConfig, error)
	}
}
type interpMaxInFlightConfigVar interp.Var

func (v MaxInFlightConfig) Interpolate(interp.Resolver) (MaxInFlightConfig, error) {
	return MaxInFlightConfig(v), nil
}
func (v MaxInFlightConfig) Wrap() interpMaxInFlightConfig {
	return interpMaxInFlightConfig{I: v}
}
func (v MaxInFlightConfig) WrapPtr() *interpMaxInFlightConfig {
	return &interpMaxInFlightConfig{I: v}
}
func (v interpMaxInFlightConfigVar) Interpolate(r interp.Resolver) (MaxInFlightConfig, error) {
	var dst MaxInFlightConfig
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpMaxInFlightConfig) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpMaxInFlightConfig) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpMaxInFlightConfigVar(v)
		return nil
	}
	var raw MaxInFlightConfig
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = MaxInFlightConfig(raw)
	return nil
}

type interpMemoryLimit struct {
	I interface {
		Interpolate(interp.Resolver) (MemoryLimit, error)
	}
}
type interpMemoryLimitVar interp.Var

func (v MemoryLimit) Interpolate(interp.Resolver) (MemoryLimit, error) {
	return MemoryLimit(v), nil
}
func (v MemoryLimit) Wrap() interpMemoryLimit     { return interpMemoryLimit{I: v} }
func (v MemoryLimit) WrapPtr() *interpMemoryLimit { return &interpMemoryLimit{I: v} }
func (v interpMemoryLimitVar) Interpolate(r interp.Resolver) (MemoryLimit, error) {
	var dst MemoryLimit
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpMemoryLimit) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpMemoryLimit) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpMemoryLimitVar(v)
		return nil
	}
	var raw MemoryLimit
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = MemoryLimit(raw)
	return nil
}

type interpVersion struct {
	I interface {
		Interpolate(interp.Resolver) (Version, error)
	}
}
type interpVersionVar interp.Var

func (v Version) Interpolate(interp.Resolver) (Version, error) {
	return Version(v), nil
}
func (v Version) Wrap() interpVersion     { return interpVersion{I: v} }
func (v Version) WrapPtr() *interpVersion { return &interpVersion{I: v} }
func (v interpVersionVar) Interpolate(r interp.Resolver) (Version, error) {
	var dst Version
	err := interp.Var(v).InterpolateInto(r, &dst)
	return dst, err
}

func (i interpVersion) MarshalJSON() ([]byte, error) { return json.Marshal(i.I) }
func (i *interpVersion) UnmarshalJSON(data []byte) error {
	var v interp.Var
	if err := json.Unmarshal(data, &v); err == nil {
		i.I = interpVersionVar(v)
		return nil
	}
	var raw Version
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	i.I = Version(raw)
	return nil
}
