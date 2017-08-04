// Generated from ../FoxySheep.g4 by ANTLR 4.7.

package parser // FoxySheep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FoxySheepListener is a complete listener for a parse tree produced by FoxySheepParser.
type FoxySheepListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterPatternExp is called when entering the PatternExp production.
	EnterPatternExp(c *PatternExpContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterConjugate is called when entering the Conjugate production.
	EnterConjugate(c *ConjugateContext)

	// EnterCeiling is called when entering the Ceiling production.
	EnterCeiling(c *CeilingContext)

	// EnterInfix is called when entering the Infix production.
	EnterInfix(c *InfixContext)

	// EnterTherefore is called when entering the Therefore production.
	EnterTherefore(c *ThereforeContext)

	// EnterTagUnset is called when entering the TagUnset production.
	EnterTagUnset(c *TagUnsetContext)

	// EnterAccessor is called when entering the Accessor production.
	EnterAccessor(c *AccessorContext)

	// EnterCircleMinus is called when entering the CircleMinus production.
	EnterCircleMinus(c *CircleMinusContext)

	// EnterDivide is called when entering the Divide production.
	EnterDivide(c *DivideContext)

	// EnterImplies is called when entering the Implies production.
	EnterImplies(c *ImpliesContext)

	// EnterPlusOp is called when entering the PlusOp production.
	EnterPlusOp(c *PlusOpContext)

	// EnterRightComposition is called when entering the RightComposition production.
	EnterRightComposition(c *RightCompositionContext)

	// EnterNonCommutativeMultiply is called when entering the NonCommutativeMultiply production.
	EnterNonCommutativeMultiply(c *NonCommutativeMultiplyContext)

	// EnterList is called when entering the List production.
	EnterList(c *ListContext)

	// EnterCup is called when entering the Cup production.
	EnterCup(c *CupContext)

	// EnterSame is called when entering the Same production.
	EnterSame(c *SameContext)

	// EnterBoxParen is called when entering the BoxParen production.
	EnterBoxParen(c *BoxParenContext)

	// EnterOptional is called when entering the Optional production.
	EnterOptional(c *OptionalContext)

	// EnterSuchThat is called when entering the SuchThat production.
	EnterSuchThat(c *SuchThatContext)

	// EnterDoubleBracketingBar is called when entering the DoubleBracketingBar production.
	EnterDoubleBracketingBar(c *DoubleBracketingBarContext)

	// EnterPatternBlankDot is called when entering the PatternBlankDot production.
	EnterPatternBlankDot(c *PatternBlankDotContext)

	// EnterDot is called when entering the Dot production.
	EnterDot(c *DotContext)

	// EnterVerticalBar is called when entering the VerticalBar production.
	EnterVerticalBar(c *VerticalBarContext)

	// EnterSquare is called when entering the Square production.
	EnterSquare(c *SquareContext)

	// EnterAlternatives is called when entering the Alternatives production.
	EnterAlternatives(c *AlternativesContext)

	// EnterOut is called when entering the Out production.
	EnterOut(c *OutContext)

	// EnterBoxConstructor is called when entering the BoxConstructor production.
	EnterBoxConstructor(c *BoxConstructorContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterPostfix is called when entering the Postfix production.
	EnterPostfix(c *PostfixContext)

	// EnterPatternBlanks is called when entering the PatternBlanks production.
	EnterPatternBlanks(c *PatternBlanksContext)

	// EnterUnaryPlusMinus is called when entering the UnaryPlusMinus production.
	EnterUnaryPlusMinus(c *UnaryPlusMinusContext)

	// EnterCap is called when entering the Cap production.
	EnterCap(c *CapContext)

	// EnterCirclePlus is called when entering the CirclePlus production.
	EnterCirclePlus(c *CirclePlusContext)

	// EnterBecause is called when entering the Because production.
	EnterBecause(c *BecauseContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// EnterGet is called when entering the Get production.
	EnterGet(c *GetContext)

	// EnterEquivalent is called when entering the Equivalent production.
	EnterEquivalent(c *EquivalentContext)

	// EnterCompoundExpression is called when entering the CompoundExpression production.
	EnterCompoundExpression(c *CompoundExpressionContext)

	// EnterDerivative is called when entering the Derivative production.
	EnterDerivative(c *DerivativeContext)

	// EnterSlot is called when entering the Slot production.
	EnterSlot(c *SlotContext)

	// EnterRightTee is called when entering the RightTee production.
	EnterRightTee(c *RightTeeContext)

	// EnterXor is called when entering the Xor production.
	EnterXor(c *XorContext)

	// EnterRule is called when entering the Rule production.
	EnterRule(c *RuleContext)

	// EnterHeadExpression is called when entering the HeadExpression production.
	EnterHeadExpression(c *HeadExpressionContext)

	// EnterReplaceAll is called when entering the ReplaceAll production.
	EnterReplaceAll(c *ReplaceAllContext)

	// EnterIntersection is called when entering the Intersection production.
	EnterIntersection(c *IntersectionContext)

	// EnterPreIncrement is called when entering the PreIncrement production.
	EnterPreIncrement(c *PreIncrementContext)

	// EnterIntegrate is called when entering the Integrate production.
	EnterIntegrate(c *IntegrateContext)

	// EnterSet is called when entering the Set production.
	EnterSet(c *SetContext)

	// EnterOpEquals is called when entering the OpEquals production.
	EnterOpEquals(c *OpEqualsContext)

	// EnterMessage is called when entering the Message production.
	EnterMessage(c *MessageContext)

	// EnterCross is called when entering the Cross production.
	EnterCross(c *CrossContext)

	// EnterPatternTest is called when entering the PatternTest production.
	EnterPatternTest(c *PatternTestContext)

	// EnterPrefix is called when entering the Prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterBackslash is called when entering the Backslash production.
	EnterBackslash(c *BackslashContext)

	// EnterRepeated is called when entering the Repeated production.
	EnterRepeated(c *RepeatedContext)

	// EnterMapApply is called when entering the MapApply production.
	EnterMapApply(c *MapApplyContext)

	// EnterUnion is called when entering the Union production.
	EnterUnion(c *UnionContext)

	// EnterVerticalSeparator is called when entering the VerticalSeparator production.
	EnterVerticalSeparator(c *VerticalSeparatorContext)

	// EnterFactorial is called when entering the Factorial production.
	EnterFactorial(c *FactorialContext)

	// EnterSpanA is called when entering the SpanA production.
	EnterSpanA(c *SpanAContext)

	// EnterFunction is called when entering the Function production.
	EnterFunction(c *FunctionContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterStar is called when entering the Star production.
	EnterStar(c *StarContext)

	// EnterComparison is called when entering the Comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterTagSet is called when entering the TagSet production.
	EnterTagSet(c *TagSetContext)

	// EnterIncrement is called when entering the Increment production.
	EnterIncrement(c *IncrementContext)

	// EnterVerticalTilde is called when entering the VerticalTilde production.
	EnterVerticalTilde(c *VerticalTildeContext)

	// EnterColon is called when entering the Colon production.
	EnterColon(c *ColonContext)

	// EnterSmallCircle is called when entering the SmallCircle production.
	EnterSmallCircle(c *SmallCircleContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterSpanB is called when entering the SpanB production.
	EnterSpanB(c *SpanBContext)

	// EnterCondition is called when entering the Condition production.
	EnterCondition(c *ConditionContext)

	// EnterFloor is called when entering the Floor production.
	EnterFloor(c *FloorContext)

	// EnterComposition is called when entering the Composition production.
	EnterComposition(c *CompositionContext)

	// EnterCircleDot is called when entering the CircleDot production.
	EnterCircleDot(c *CircleDotContext)

	// EnterSymbolLiteral is called when entering the SymbolLiteral production.
	EnterSymbolLiteral(c *SymbolLiteralContext)

	// EnterCircleTimes is called when entering the CircleTimes production.
	EnterCircleTimes(c *CircleTimesContext)

	// EnterUnset is called when entering the Unset production.
	EnterUnset(c *UnsetContext)

	// EnterDel is called when entering the Del production.
	EnterDel(c *DelContext)

	// EnterDiamond is called when entering the Diamond production.
	EnterDiamond(c *DiamondContext)

	// EnterWedge is called when entering the Wedge production.
	EnterWedge(c *WedgeContext)

	// EnterPut is called when entering the Put production.
	EnterPut(c *PutContext)

	// EnterStringJoin is called when entering the StringJoin production.
	EnterStringJoin(c *StringJoinContext)

	// EnterTee is called when entering the Tee production.
	EnterTee(c *TeeContext)

	// EnterSetContainment is called when entering the SetContainment production.
	EnterSetContainment(c *SetContainmentContext)

	// EnterVee is called when entering the Vee production.
	EnterVee(c *VeeContext)

	// EnterCenterDot is called when entering the CenterDot production.
	EnterCenterDot(c *CenterDotContext)

	// EnterTimes is called when entering the Times production.
	EnterTimes(c *TimesContext)

	// EnterStringExpression is called when entering the StringExpression production.
	EnterStringExpression(c *StringExpressionContext)

	// EnterBracketingBar is called when entering the BracketingBar production.
	EnterBracketingBar(c *BracketingBarContext)

	// EnterCoproduct is called when entering the Coproduct production.
	EnterCoproduct(c *CoproductContext)

	// EnterAngleBracket is called when entering the AngleBracket production.
	EnterAngleBracket(c *AngleBracketContext)

	// EnterPower is called when entering the Power production.
	EnterPower(c *PowerContext)

	// EnterContextName is called when entering the ContextName production.
	EnterContextName(c *ContextNameContext)

	// EnterSimpleContext is called when entering the SimpleContext production.
	EnterSimpleContext(c *SimpleContextContext)

	// EnterCompoundContext is called when entering the CompoundContext production.
	EnterCompoundContext(c *CompoundContextContext)

	// EnterNumberBaseN is called when entering the NumberBaseN production.
	EnterNumberBaseN(c *NumberBaseNContext)

	// EnterNumberBaseTen is called when entering the NumberBaseTen production.
	EnterNumberBaseTen(c *NumberBaseTenContext)

	// EnterNumberLiteralPrecision is called when entering the numberLiteralPrecision production.
	EnterNumberLiteralPrecision(c *NumberLiteralPrecisionContext)

	// EnterNumberLiteralExponent is called when entering the numberLiteralExponent production.
	EnterNumberLiteralExponent(c *NumberLiteralExponentContext)

	// EnterOutNumbered is called when entering the OutNumbered production.
	EnterOutNumbered(c *OutNumberedContext)

	// EnterOutUnnumbered is called when entering the OutUnnumbered production.
	EnterOutUnnumbered(c *OutUnnumberedContext)

	// EnterSlotDigits is called when entering the SlotDigits production.
	EnterSlotDigits(c *SlotDigitsContext)

	// EnterSlotNamed is called when entering the SlotNamed production.
	EnterSlotNamed(c *SlotNamedContext)

	// EnterSlotSequenceDigits is called when entering the SlotSequenceDigits production.
	EnterSlotSequenceDigits(c *SlotSequenceDigitsContext)

	// EnterSlotSequence is called when entering the SlotSequence production.
	EnterSlotSequence(c *SlotSequenceContext)

	// EnterSlotUnnamed is called when entering the SlotUnnamed production.
	EnterSlotUnnamed(c *SlotUnnamedContext)

	// EnterExpressionListed is called when entering the ExpressionListed production.
	EnterExpressionListed(c *ExpressionListedContext)

	// EnterAccessExpressionA is called when entering the AccessExpressionA production.
	EnterAccessExpressionA(c *AccessExpressionAContext)

	// EnterAccessExpressionB is called when entering the AccessExpressionB production.
	EnterAccessExpressionB(c *AccessExpressionBContext)

	// EnterBox is called when entering the box production.
	EnterBox(c *BoxContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitPatternExp is called when exiting the PatternExp production.
	ExitPatternExp(c *PatternExpContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitConjugate is called when exiting the Conjugate production.
	ExitConjugate(c *ConjugateContext)

	// ExitCeiling is called when exiting the Ceiling production.
	ExitCeiling(c *CeilingContext)

	// ExitInfix is called when exiting the Infix production.
	ExitInfix(c *InfixContext)

	// ExitTherefore is called when exiting the Therefore production.
	ExitTherefore(c *ThereforeContext)

	// ExitTagUnset is called when exiting the TagUnset production.
	ExitTagUnset(c *TagUnsetContext)

	// ExitAccessor is called when exiting the Accessor production.
	ExitAccessor(c *AccessorContext)

	// ExitCircleMinus is called when exiting the CircleMinus production.
	ExitCircleMinus(c *CircleMinusContext)

	// ExitDivide is called when exiting the Divide production.
	ExitDivide(c *DivideContext)

	// ExitImplies is called when exiting the Implies production.
	ExitImplies(c *ImpliesContext)

	// ExitPlusOp is called when exiting the PlusOp production.
	ExitPlusOp(c *PlusOpContext)

	// ExitRightComposition is called when exiting the RightComposition production.
	ExitRightComposition(c *RightCompositionContext)

	// ExitNonCommutativeMultiply is called when exiting the NonCommutativeMultiply production.
	ExitNonCommutativeMultiply(c *NonCommutativeMultiplyContext)

	// ExitList is called when exiting the List production.
	ExitList(c *ListContext)

	// ExitCup is called when exiting the Cup production.
	ExitCup(c *CupContext)

	// ExitSame is called when exiting the Same production.
	ExitSame(c *SameContext)

	// ExitBoxParen is called when exiting the BoxParen production.
	ExitBoxParen(c *BoxParenContext)

	// ExitOptional is called when exiting the Optional production.
	ExitOptional(c *OptionalContext)

	// ExitSuchThat is called when exiting the SuchThat production.
	ExitSuchThat(c *SuchThatContext)

	// ExitDoubleBracketingBar is called when exiting the DoubleBracketingBar production.
	ExitDoubleBracketingBar(c *DoubleBracketingBarContext)

	// ExitPatternBlankDot is called when exiting the PatternBlankDot production.
	ExitPatternBlankDot(c *PatternBlankDotContext)

	// ExitDot is called when exiting the Dot production.
	ExitDot(c *DotContext)

	// ExitVerticalBar is called when exiting the VerticalBar production.
	ExitVerticalBar(c *VerticalBarContext)

	// ExitSquare is called when exiting the Square production.
	ExitSquare(c *SquareContext)

	// ExitAlternatives is called when exiting the Alternatives production.
	ExitAlternatives(c *AlternativesContext)

	// ExitOut is called when exiting the Out production.
	ExitOut(c *OutContext)

	// ExitBoxConstructor is called when exiting the BoxConstructor production.
	ExitBoxConstructor(c *BoxConstructorContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitPostfix is called when exiting the Postfix production.
	ExitPostfix(c *PostfixContext)

	// ExitPatternBlanks is called when exiting the PatternBlanks production.
	ExitPatternBlanks(c *PatternBlanksContext)

	// ExitUnaryPlusMinus is called when exiting the UnaryPlusMinus production.
	ExitUnaryPlusMinus(c *UnaryPlusMinusContext)

	// ExitCap is called when exiting the Cap production.
	ExitCap(c *CapContext)

	// ExitCirclePlus is called when exiting the CirclePlus production.
	ExitCirclePlus(c *CirclePlusContext)

	// ExitBecause is called when exiting the Because production.
	ExitBecause(c *BecauseContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)

	// ExitGet is called when exiting the Get production.
	ExitGet(c *GetContext)

	// ExitEquivalent is called when exiting the Equivalent production.
	ExitEquivalent(c *EquivalentContext)

	// ExitCompoundExpression is called when exiting the CompoundExpression production.
	ExitCompoundExpression(c *CompoundExpressionContext)

	// ExitDerivative is called when exiting the Derivative production.
	ExitDerivative(c *DerivativeContext)

	// ExitSlot is called when exiting the Slot production.
	ExitSlot(c *SlotContext)

	// ExitRightTee is called when exiting the RightTee production.
	ExitRightTee(c *RightTeeContext)

	// ExitXor is called when exiting the Xor production.
	ExitXor(c *XorContext)

	// ExitRule is called when exiting the Rule production.
	ExitRule(c *RuleContext)

	// ExitHeadExpression is called when exiting the HeadExpression production.
	ExitHeadExpression(c *HeadExpressionContext)

	// ExitReplaceAll is called when exiting the ReplaceAll production.
	ExitReplaceAll(c *ReplaceAllContext)

	// ExitIntersection is called when exiting the Intersection production.
	ExitIntersection(c *IntersectionContext)

	// ExitPreIncrement is called when exiting the PreIncrement production.
	ExitPreIncrement(c *PreIncrementContext)

	// ExitIntegrate is called when exiting the Integrate production.
	ExitIntegrate(c *IntegrateContext)

	// ExitSet is called when exiting the Set production.
	ExitSet(c *SetContext)

	// ExitOpEquals is called when exiting the OpEquals production.
	ExitOpEquals(c *OpEqualsContext)

	// ExitMessage is called when exiting the Message production.
	ExitMessage(c *MessageContext)

	// ExitCross is called when exiting the Cross production.
	ExitCross(c *CrossContext)

	// ExitPatternTest is called when exiting the PatternTest production.
	ExitPatternTest(c *PatternTestContext)

	// ExitPrefix is called when exiting the Prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitBackslash is called when exiting the Backslash production.
	ExitBackslash(c *BackslashContext)

	// ExitRepeated is called when exiting the Repeated production.
	ExitRepeated(c *RepeatedContext)

	// ExitMapApply is called when exiting the MapApply production.
	ExitMapApply(c *MapApplyContext)

	// ExitUnion is called when exiting the Union production.
	ExitUnion(c *UnionContext)

	// ExitVerticalSeparator is called when exiting the VerticalSeparator production.
	ExitVerticalSeparator(c *VerticalSeparatorContext)

	// ExitFactorial is called when exiting the Factorial production.
	ExitFactorial(c *FactorialContext)

	// ExitSpanA is called when exiting the SpanA production.
	ExitSpanA(c *SpanAContext)

	// ExitFunction is called when exiting the Function production.
	ExitFunction(c *FunctionContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitStar is called when exiting the Star production.
	ExitStar(c *StarContext)

	// ExitComparison is called when exiting the Comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitTagSet is called when exiting the TagSet production.
	ExitTagSet(c *TagSetContext)

	// ExitIncrement is called when exiting the Increment production.
	ExitIncrement(c *IncrementContext)

	// ExitVerticalTilde is called when exiting the VerticalTilde production.
	ExitVerticalTilde(c *VerticalTildeContext)

	// ExitColon is called when exiting the Colon production.
	ExitColon(c *ColonContext)

	// ExitSmallCircle is called when exiting the SmallCircle production.
	ExitSmallCircle(c *SmallCircleContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitSpanB is called when exiting the SpanB production.
	ExitSpanB(c *SpanBContext)

	// ExitCondition is called when exiting the Condition production.
	ExitCondition(c *ConditionContext)

	// ExitFloor is called when exiting the Floor production.
	ExitFloor(c *FloorContext)

	// ExitComposition is called when exiting the Composition production.
	ExitComposition(c *CompositionContext)

	// ExitCircleDot is called when exiting the CircleDot production.
	ExitCircleDot(c *CircleDotContext)

	// ExitSymbolLiteral is called when exiting the SymbolLiteral production.
	ExitSymbolLiteral(c *SymbolLiteralContext)

	// ExitCircleTimes is called when exiting the CircleTimes production.
	ExitCircleTimes(c *CircleTimesContext)

	// ExitUnset is called when exiting the Unset production.
	ExitUnset(c *UnsetContext)

	// ExitDel is called when exiting the Del production.
	ExitDel(c *DelContext)

	// ExitDiamond is called when exiting the Diamond production.
	ExitDiamond(c *DiamondContext)

	// ExitWedge is called when exiting the Wedge production.
	ExitWedge(c *WedgeContext)

	// ExitPut is called when exiting the Put production.
	ExitPut(c *PutContext)

	// ExitStringJoin is called when exiting the StringJoin production.
	ExitStringJoin(c *StringJoinContext)

	// ExitTee is called when exiting the Tee production.
	ExitTee(c *TeeContext)

	// ExitSetContainment is called when exiting the SetContainment production.
	ExitSetContainment(c *SetContainmentContext)

	// ExitVee is called when exiting the Vee production.
	ExitVee(c *VeeContext)

	// ExitCenterDot is called when exiting the CenterDot production.
	ExitCenterDot(c *CenterDotContext)

	// ExitTimes is called when exiting the Times production.
	ExitTimes(c *TimesContext)

	// ExitStringExpression is called when exiting the StringExpression production.
	ExitStringExpression(c *StringExpressionContext)

	// ExitBracketingBar is called when exiting the BracketingBar production.
	ExitBracketingBar(c *BracketingBarContext)

	// ExitCoproduct is called when exiting the Coproduct production.
	ExitCoproduct(c *CoproductContext)

	// ExitAngleBracket is called when exiting the AngleBracket production.
	ExitAngleBracket(c *AngleBracketContext)

	// ExitPower is called when exiting the Power production.
	ExitPower(c *PowerContext)

	// ExitContextName is called when exiting the ContextName production.
	ExitContextName(c *ContextNameContext)

	// ExitSimpleContext is called when exiting the SimpleContext production.
	ExitSimpleContext(c *SimpleContextContext)

	// ExitCompoundContext is called when exiting the CompoundContext production.
	ExitCompoundContext(c *CompoundContextContext)

	// ExitNumberBaseN is called when exiting the NumberBaseN production.
	ExitNumberBaseN(c *NumberBaseNContext)

	// ExitNumberBaseTen is called when exiting the NumberBaseTen production.
	ExitNumberBaseTen(c *NumberBaseTenContext)

	// ExitNumberLiteralPrecision is called when exiting the numberLiteralPrecision production.
	ExitNumberLiteralPrecision(c *NumberLiteralPrecisionContext)

	// ExitNumberLiteralExponent is called when exiting the numberLiteralExponent production.
	ExitNumberLiteralExponent(c *NumberLiteralExponentContext)

	// ExitOutNumbered is called when exiting the OutNumbered production.
	ExitOutNumbered(c *OutNumberedContext)

	// ExitOutUnnumbered is called when exiting the OutUnnumbered production.
	ExitOutUnnumbered(c *OutUnnumberedContext)

	// ExitSlotDigits is called when exiting the SlotDigits production.
	ExitSlotDigits(c *SlotDigitsContext)

	// ExitSlotNamed is called when exiting the SlotNamed production.
	ExitSlotNamed(c *SlotNamedContext)

	// ExitSlotSequenceDigits is called when exiting the SlotSequenceDigits production.
	ExitSlotSequenceDigits(c *SlotSequenceDigitsContext)

	// ExitSlotSequence is called when exiting the SlotSequence production.
	ExitSlotSequence(c *SlotSequenceContext)

	// ExitSlotUnnamed is called when exiting the SlotUnnamed production.
	ExitSlotUnnamed(c *SlotUnnamedContext)

	// ExitExpressionListed is called when exiting the ExpressionListed production.
	ExitExpressionListed(c *ExpressionListedContext)

	// ExitAccessExpressionA is called when exiting the AccessExpressionA production.
	ExitAccessExpressionA(c *AccessExpressionAContext)

	// ExitAccessExpressionB is called when exiting the AccessExpressionB production.
	ExitAccessExpressionB(c *AccessExpressionBContext)

	// ExitBox is called when exiting the box production.
	ExitBox(c *BoxContext)
}
