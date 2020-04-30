// Code generated by "stringer -output=pkg/sql/opt/rule_name_string.go -type=RuleName pkg/sql/opt/rule_name.go pkg/sql/opt/rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidRuleName-0]
	_ = x[SimplifyRootOrdering-1]
	_ = x[PruneRootCols-2]
	_ = x[SimplifyZeroCardinalityGroup-3]
	_ = x[NumManualRuleNames-4]
	_ = x[startAutoRule-4]
	_ = x[EliminateAggDistinct-5]
	_ = x[NormalizeNestedAnds-6]
	_ = x[SimplifyTrueAnd-7]
	_ = x[SimplifyAndTrue-8]
	_ = x[SimplifyFalseAnd-9]
	_ = x[SimplifyAndFalse-10]
	_ = x[SimplifyTrueOr-11]
	_ = x[SimplifyOrTrue-12]
	_ = x[SimplifyFalseOr-13]
	_ = x[SimplifyOrFalse-14]
	_ = x[SimplifyRange-15]
	_ = x[FoldNullAndOr-16]
	_ = x[FoldNotTrue-17]
	_ = x[FoldNotFalse-18]
	_ = x[FoldNotNull-19]
	_ = x[NegateComparison-20]
	_ = x[EliminateNot-21]
	_ = x[NegateAnd-22]
	_ = x[NegateOr-23]
	_ = x[ExtractRedundantConjunct-24]
	_ = x[CommuteVarInequality-25]
	_ = x[CommuteConstInequality-26]
	_ = x[NormalizeCmpPlusConst-27]
	_ = x[NormalizeCmpMinusConst-28]
	_ = x[NormalizeCmpConstMinus-29]
	_ = x[NormalizeTupleEquality-30]
	_ = x[FoldNullComparisonLeft-31]
	_ = x[FoldNullComparisonRight-32]
	_ = x[FoldIsNull-33]
	_ = x[FoldNonNullIsNull-34]
	_ = x[FoldIsNotNull-35]
	_ = x[FoldNonNullIsNotNull-36]
	_ = x[CommuteNullIs-37]
	_ = x[DecorrelateJoin-38]
	_ = x[DecorrelateProjectSet-39]
	_ = x[TryDecorrelateSelect-40]
	_ = x[TryDecorrelateProject-41]
	_ = x[TryDecorrelateProjectSelect-42]
	_ = x[TryDecorrelateProjectInnerJoin-43]
	_ = x[TryDecorrelateInnerJoin-44]
	_ = x[TryDecorrelateInnerLeftJoin-45]
	_ = x[TryDecorrelateGroupBy-46]
	_ = x[TryDecorrelateScalarGroupBy-47]
	_ = x[TryDecorrelateSemiJoin-48]
	_ = x[TryDecorrelateLimitOne-49]
	_ = x[TryDecorrelateProjectSet-50]
	_ = x[TryDecorrelateWindow-51]
	_ = x[TryDecorrelateMax1Row-52]
	_ = x[HoistSelectExists-53]
	_ = x[HoistSelectNotExists-54]
	_ = x[HoistSelectSubquery-55]
	_ = x[HoistProjectSubquery-56]
	_ = x[HoistJoinSubquery-57]
	_ = x[HoistValuesSubquery-58]
	_ = x[HoistProjectSetSubquery-59]
	_ = x[NormalizeSelectAnyFilter-60]
	_ = x[NormalizeJoinAnyFilter-61]
	_ = x[NormalizeSelectNotAnyFilter-62]
	_ = x[NormalizeJoinNotAnyFilter-63]
	_ = x[FoldNullCast-64]
	_ = x[FoldNullUnary-65]
	_ = x[FoldNullBinaryLeft-66]
	_ = x[FoldNullBinaryRight-67]
	_ = x[FoldNullInNonEmpty-68]
	_ = x[FoldInEmpty-69]
	_ = x[FoldNotInEmpty-70]
	_ = x[FoldArray-71]
	_ = x[FoldBinary-72]
	_ = x[FoldUnary-73]
	_ = x[FoldComparison-74]
	_ = x[FoldCast-75]
	_ = x[FoldIndirection-76]
	_ = x[FoldColumnAccess-77]
	_ = x[FoldFunction-78]
	_ = x[FoldEqualsAnyNull-79]
	_ = x[ConvertGroupByToDistinct-80]
	_ = x[EliminateDistinct-81]
	_ = x[EliminateGroupByProject-82]
	_ = x[ReduceGroupingCols-83]
	_ = x[ReduceNotNullGroupingCols-84]
	_ = x[EliminateAggDistinctForKeys-85]
	_ = x[EliminateAggFilteredDistinctForKeys-86]
	_ = x[EliminateDistinctNoColumns-87]
	_ = x[EliminateEnsureDistinctNoColumns-88]
	_ = x[EliminateDistinctOnValues-89]
	_ = x[PushAggDistinctIntoScalarGroupBy-90]
	_ = x[PushAggFilterIntoScalarGroupBy-91]
	_ = x[ConvertCountToCountRows-92]
	_ = x[InlineProjectConstants-93]
	_ = x[InlineSelectConstants-94]
	_ = x[InlineJoinConstantsLeft-95]
	_ = x[InlineJoinConstantsRight-96]
	_ = x[PushSelectIntoInlinableProject-97]
	_ = x[InlineProjectInProject-98]
	_ = x[CommuteRightJoin-99]
	_ = x[SimplifyJoinFilters-100]
	_ = x[DetectJoinContradiction-101]
	_ = x[PushFilterIntoJoinLeftAndRight-102]
	_ = x[MapFilterIntoJoinLeft-103]
	_ = x[MapFilterIntoJoinRight-104]
	_ = x[MapEqualityIntoJoinLeftAndRight-105]
	_ = x[PushFilterIntoJoinLeft-106]
	_ = x[PushFilterIntoJoinRight-107]
	_ = x[SimplifyLeftJoinWithoutFilters-108]
	_ = x[SimplifyRightJoinWithoutFilters-109]
	_ = x[SimplifyLeftJoinWithFilters-110]
	_ = x[SimplifyRightJoinWithFilters-111]
	_ = x[EliminateSemiJoin-112]
	_ = x[SimplifyZeroCardinalitySemiJoin-113]
	_ = x[EliminateAntiJoin-114]
	_ = x[SimplifyZeroCardinalityAntiJoin-115]
	_ = x[EliminateJoinNoColsLeft-116]
	_ = x[EliminateJoinNoColsRight-117]
	_ = x[HoistJoinProjectRight-118]
	_ = x[HoistJoinProjectLeft-119]
	_ = x[SimplifyJoinNotNullEquality-120]
	_ = x[ExtractJoinEqualities-121]
	_ = x[SortFiltersInJoin-122]
	_ = x[EliminateLimit-123]
	_ = x[EliminateOffset-124]
	_ = x[PushLimitIntoProject-125]
	_ = x[PushOffsetIntoProject-126]
	_ = x[PushLimitIntoOffset-127]
	_ = x[PushLimitIntoOrdinality-128]
	_ = x[PushLimitIntoLeftJoin-129]
	_ = x[EliminateMax1Row-130]
	_ = x[FoldPlusZero-131]
	_ = x[FoldZeroPlus-132]
	_ = x[FoldMinusZero-133]
	_ = x[FoldMultOne-134]
	_ = x[FoldOneMult-135]
	_ = x[FoldDivOne-136]
	_ = x[InvertMinus-137]
	_ = x[EliminateUnaryMinus-138]
	_ = x[SimplifyLimitOrdering-139]
	_ = x[SimplifyOffsetOrdering-140]
	_ = x[SimplifyGroupByOrdering-141]
	_ = x[SimplifyOrdinalityOrdering-142]
	_ = x[SimplifyExplainOrdering-143]
	_ = x[EliminateProject-144]
	_ = x[MergeProjects-145]
	_ = x[MergeProjectWithValues-146]
	_ = x[ConvertZipArraysToValues-147]
	_ = x[PruneProjectCols-148]
	_ = x[PruneScanCols-149]
	_ = x[PruneSelectCols-150]
	_ = x[PruneLimitCols-151]
	_ = x[PruneOffsetCols-152]
	_ = x[PruneJoinLeftCols-153]
	_ = x[PruneJoinRightCols-154]
	_ = x[PruneSemiAntiJoinRightCols-155]
	_ = x[PruneAggCols-156]
	_ = x[PruneGroupByCols-157]
	_ = x[PruneValuesCols-158]
	_ = x[PruneOrdinalityCols-159]
	_ = x[PruneExplainCols-160]
	_ = x[PruneProjectSetCols-161]
	_ = x[PruneWindowOutputCols-162]
	_ = x[PruneWindowInputCols-163]
	_ = x[PruneMutationFetchCols-164]
	_ = x[PruneMutationInputCols-165]
	_ = x[PruneMutationReturnCols-166]
	_ = x[PruneWithScanCols-167]
	_ = x[PruneWithCols-168]
	_ = x[PruneUnionAllCols-169]
	_ = x[RejectNullsLeftJoin-170]
	_ = x[RejectNullsRightJoin-171]
	_ = x[RejectNullsGroupBy-172]
	_ = x[CommuteVar-173]
	_ = x[CommuteConst-174]
	_ = x[EliminateCoalesce-175]
	_ = x[SimplifyCoalesce-176]
	_ = x[EliminateCast-177]
	_ = x[NormalizeInConst-178]
	_ = x[FoldInNull-179]
	_ = x[UnifyComparisonTypes-180]
	_ = x[EliminateExistsZeroRows-181]
	_ = x[EliminateExistsProject-182]
	_ = x[EliminateExistsGroupBy-183]
	_ = x[IntroduceExistsLimit-184]
	_ = x[EliminateExistsLimit-185]
	_ = x[NormalizeJSONFieldAccess-186]
	_ = x[NormalizeJSONContains-187]
	_ = x[SimplifyCaseWhenConstValue-188]
	_ = x[InlineAnyValuesSingleCol-189]
	_ = x[InlineAnyValuesMultiCol-190]
	_ = x[SimplifyEqualsAnyTuple-191]
	_ = x[SimplifyAnyScalarArray-192]
	_ = x[FoldCollate-193]
	_ = x[NormalizeArrayFlattenToAgg-194]
	_ = x[SimplifySameVarEqualities-195]
	_ = x[SimplifySameVarInequalities-196]
	_ = x[SimplifySelectFilters-197]
	_ = x[ConsolidateSelectFilters-198]
	_ = x[DetectSelectContradiction-199]
	_ = x[EliminateSelect-200]
	_ = x[MergeSelects-201]
	_ = x[PushSelectIntoProject-202]
	_ = x[MergeSelectInnerJoin-203]
	_ = x[PushSelectCondLeftIntoJoinLeftAndRight-204]
	_ = x[PushSelectIntoJoinLeft-205]
	_ = x[PushSelectIntoGroupBy-206]
	_ = x[RemoveNotNullCondition-207]
	_ = x[InlineConstVar-208]
	_ = x[PushSelectIntoProjectSet-209]
	_ = x[PushFilterIntoSetOp-210]
	_ = x[EliminateUnionAllLeft-211]
	_ = x[EliminateUnionAllRight-212]
	_ = x[EliminateWindow-213]
	_ = x[ReduceWindowPartitionCols-214]
	_ = x[SimplifyWindowOrdering-215]
	_ = x[PushSelectIntoWindow-216]
	_ = x[PushLimitIntoWindow-217]
	_ = x[InlineWith-218]
	_ = x[startExploreRule-219]
	_ = x[ReplaceScalarMinMaxWithLimit-220]
	_ = x[ReplaceMinWithLimit-221]
	_ = x[ReplaceMaxWithLimit-222]
	_ = x[GenerateStreamingGroupBy-223]
	_ = x[CommuteJoin-224]
	_ = x[CommuteLeftJoin-225]
	_ = x[CommuteSemiJoin-226]
	_ = x[GenerateMergeJoins-227]
	_ = x[GenerateLookupJoins-228]
	_ = x[GenerateGeoLookupJoins-229]
	_ = x[GenerateZigzagJoins-230]
	_ = x[GenerateInvertedIndexZigzagJoins-231]
	_ = x[GenerateLookupJoinsWithFilter-232]
	_ = x[AssociateJoin-233]
	_ = x[GenerateLimitedScans-234]
	_ = x[PushLimitIntoConstrainedScan-235]
	_ = x[PushLimitIntoIndexJoin-236]
	_ = x[GenerateIndexScans-237]
	_ = x[GenerateConstrainedScans-238]
	_ = x[GenerateInvertedIndexScans-239]
	_ = x[SplitDisjunction-240]
	_ = x[SplitDisjunctionAddKey-241]
	_ = x[NumRuleNames-242]
}

const _RuleName_name = "InvalidRuleNameSimplifyRootOrderingPruneRootColsSimplifyZeroCardinalityGroupNumManualRuleNamesEliminateAggDistinctNormalizeNestedAndsSimplifyTrueAndSimplifyAndTrueSimplifyFalseAndSimplifyAndFalseSimplifyTrueOrSimplifyOrTrueSimplifyFalseOrSimplifyOrFalseSimplifyRangeFoldNullAndOrFoldNotTrueFoldNotFalseFoldNotNullNegateComparisonEliminateNotNegateAndNegateOrExtractRedundantConjunctCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightFoldIsNullFoldNonNullIsNullFoldIsNotNullFoldNonNullIsNotNullCommuteNullIsDecorrelateJoinDecorrelateProjectSetTryDecorrelateSelectTryDecorrelateProjectTryDecorrelateProjectSelectTryDecorrelateProjectInnerJoinTryDecorrelateInnerJoinTryDecorrelateInnerLeftJoinTryDecorrelateGroupByTryDecorrelateScalarGroupByTryDecorrelateSemiJoinTryDecorrelateLimitOneTryDecorrelateProjectSetTryDecorrelateWindowTryDecorrelateMax1RowHoistSelectExistsHoistSelectNotExistsHoistSelectSubqueryHoistProjectSubqueryHoistJoinSubqueryHoistValuesSubqueryHoistProjectSetSubqueryNormalizeSelectAnyFilterNormalizeJoinAnyFilterNormalizeSelectNotAnyFilterNormalizeJoinNotAnyFilterFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldInEmptyFoldNotInEmptyFoldArrayFoldBinaryFoldUnaryFoldComparisonFoldCastFoldIndirectionFoldColumnAccessFoldFunctionFoldEqualsAnyNullConvertGroupByToDistinctEliminateDistinctEliminateGroupByProjectReduceGroupingColsReduceNotNullGroupingColsEliminateAggDistinctForKeysEliminateAggFilteredDistinctForKeysEliminateDistinctNoColumnsEliminateEnsureDistinctNoColumnsEliminateDistinctOnValuesPushAggDistinctIntoScalarGroupByPushAggFilterIntoScalarGroupByConvertCountToCountRowsInlineProjectConstantsInlineSelectConstantsInlineJoinConstantsLeftInlineJoinConstantsRightPushSelectIntoInlinableProjectInlineProjectInProjectCommuteRightJoinSimplifyJoinFiltersDetectJoinContradictionPushFilterIntoJoinLeftAndRightMapFilterIntoJoinLeftMapFilterIntoJoinRightMapEqualityIntoJoinLeftAndRightPushFilterIntoJoinLeftPushFilterIntoJoinRightSimplifyLeftJoinWithoutFiltersSimplifyRightJoinWithoutFiltersSimplifyLeftJoinWithFiltersSimplifyRightJoinWithFiltersEliminateSemiJoinSimplifyZeroCardinalitySemiJoinEliminateAntiJoinSimplifyZeroCardinalityAntiJoinEliminateJoinNoColsLeftEliminateJoinNoColsRightHoistJoinProjectRightHoistJoinProjectLeftSimplifyJoinNotNullEqualityExtractJoinEqualitiesSortFiltersInJoinEliminateLimitEliminateOffsetPushLimitIntoProjectPushOffsetIntoProjectPushLimitIntoOffsetPushLimitIntoOrdinalityPushLimitIntoLeftJoinEliminateMax1RowFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusSimplifyLimitOrderingSimplifyOffsetOrderingSimplifyGroupByOrderingSimplifyOrdinalityOrderingSimplifyExplainOrderingEliminateProjectMergeProjectsMergeProjectWithValuesConvertZipArraysToValuesPruneProjectColsPruneScanColsPruneSelectColsPruneLimitColsPruneOffsetColsPruneJoinLeftColsPruneJoinRightColsPruneSemiAntiJoinRightColsPruneAggColsPruneGroupByColsPruneValuesColsPruneOrdinalityColsPruneExplainColsPruneProjectSetColsPruneWindowOutputColsPruneWindowInputColsPruneMutationFetchColsPruneMutationInputColsPruneMutationReturnColsPruneWithScanColsPruneWithColsPruneUnionAllColsRejectNullsLeftJoinRejectNullsRightJoinRejectNullsGroupByCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastNormalizeInConstFoldInNullUnifyComparisonTypesEliminateExistsZeroRowsEliminateExistsProjectEliminateExistsGroupByIntroduceExistsLimitEliminateExistsLimitNormalizeJSONFieldAccessNormalizeJSONContainsSimplifyCaseWhenConstValueInlineAnyValuesSingleColInlineAnyValuesMultiColSimplifyEqualsAnyTupleSimplifyAnyScalarArrayFoldCollateNormalizeArrayFlattenToAggSimplifySameVarEqualitiesSimplifySameVarInequalitiesSimplifySelectFiltersConsolidateSelectFiltersDetectSelectContradictionEliminateSelectMergeSelectsPushSelectIntoProjectMergeSelectInnerJoinPushSelectCondLeftIntoJoinLeftAndRightPushSelectIntoJoinLeftPushSelectIntoGroupByRemoveNotNullConditionInlineConstVarPushSelectIntoProjectSetPushFilterIntoSetOpEliminateUnionAllLeftEliminateUnionAllRightEliminateWindowReduceWindowPartitionColsSimplifyWindowOrderingPushSelectIntoWindowPushLimitIntoWindowInlineWithstartExploreRuleReplaceScalarMinMaxWithLimitReplaceMinWithLimitReplaceMaxWithLimitGenerateStreamingGroupByCommuteJoinCommuteLeftJoinCommuteSemiJoinGenerateMergeJoinsGenerateLookupJoinsGenerateGeoLookupJoinsGenerateZigzagJoinsGenerateInvertedIndexZigzagJoinsGenerateLookupJoinsWithFilterAssociateJoinGenerateLimitedScansPushLimitIntoConstrainedScanPushLimitIntoIndexJoinGenerateIndexScansGenerateConstrainedScansGenerateInvertedIndexScansSplitDisjunctionSplitDisjunctionAddKeyNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 35, 48, 76, 94, 114, 133, 148, 163, 179, 195, 209, 223, 238, 253, 266, 279, 290, 302, 313, 329, 341, 350, 358, 382, 402, 424, 445, 467, 489, 511, 533, 556, 566, 583, 596, 616, 629, 644, 665, 685, 706, 733, 763, 786, 813, 834, 861, 883, 905, 929, 949, 970, 987, 1007, 1026, 1046, 1063, 1082, 1105, 1129, 1151, 1178, 1203, 1215, 1228, 1246, 1265, 1283, 1294, 1308, 1317, 1327, 1336, 1350, 1358, 1373, 1389, 1401, 1418, 1442, 1459, 1482, 1500, 1525, 1552, 1587, 1613, 1645, 1670, 1702, 1732, 1755, 1777, 1798, 1821, 1845, 1875, 1897, 1913, 1932, 1955, 1985, 2006, 2028, 2059, 2081, 2104, 2134, 2165, 2192, 2220, 2237, 2268, 2285, 2316, 2339, 2363, 2384, 2404, 2431, 2452, 2469, 2483, 2498, 2518, 2539, 2558, 2581, 2602, 2618, 2630, 2642, 2655, 2666, 2677, 2687, 2698, 2717, 2738, 2760, 2783, 2809, 2832, 2848, 2861, 2883, 2907, 2923, 2936, 2951, 2965, 2980, 2997, 3015, 3041, 3053, 3069, 3084, 3103, 3119, 3138, 3159, 3179, 3201, 3223, 3246, 3263, 3276, 3293, 3312, 3332, 3350, 3360, 3372, 3389, 3405, 3418, 3434, 3444, 3464, 3487, 3509, 3531, 3551, 3571, 3595, 3616, 3642, 3666, 3689, 3711, 3733, 3744, 3770, 3795, 3822, 3843, 3867, 3892, 3907, 3919, 3940, 3960, 3998, 4020, 4041, 4063, 4077, 4101, 4120, 4141, 4163, 4178, 4203, 4225, 4245, 4264, 4274, 4290, 4318, 4337, 4356, 4380, 4391, 4406, 4421, 4439, 4458, 4480, 4499, 4531, 4560, 4573, 4593, 4621, 4643, 4661, 4685, 4711, 4727, 4749, 4761}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}
