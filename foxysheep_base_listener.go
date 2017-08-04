// Generated from ../FoxySheep.g4 by ANTLR 4.7.

package parser // FoxySheep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFoxySheepListener is a complete listener for a parse tree produced by FoxySheepParser.
type BaseFoxySheepListener struct{}

var _ FoxySheepListener = &BaseFoxySheepListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFoxySheepListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFoxySheepListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFoxySheepListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFoxySheepListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseFoxySheepListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseFoxySheepListener) ExitProg(ctx *ProgContext) {}

// EnterPatternExp is called when production PatternExp is entered.
func (s *BaseFoxySheepListener) EnterPatternExp(ctx *PatternExpContext) {}

// ExitPatternExp is called when production PatternExp is exited.
func (s *BaseFoxySheepListener) ExitPatternExp(ctx *PatternExpContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseFoxySheepListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseFoxySheepListener) ExitOr(ctx *OrContext) {}

// EnterConjugate is called when production Conjugate is entered.
func (s *BaseFoxySheepListener) EnterConjugate(ctx *ConjugateContext) {}

// ExitConjugate is called when production Conjugate is exited.
func (s *BaseFoxySheepListener) ExitConjugate(ctx *ConjugateContext) {}

// EnterCeiling is called when production Ceiling is entered.
func (s *BaseFoxySheepListener) EnterCeiling(ctx *CeilingContext) {}

// ExitCeiling is called when production Ceiling is exited.
func (s *BaseFoxySheepListener) ExitCeiling(ctx *CeilingContext) {}

// EnterInfix is called when production Infix is entered.
func (s *BaseFoxySheepListener) EnterInfix(ctx *InfixContext) {}

// ExitInfix is called when production Infix is exited.
func (s *BaseFoxySheepListener) ExitInfix(ctx *InfixContext) {}

// EnterTherefore is called when production Therefore is entered.
func (s *BaseFoxySheepListener) EnterTherefore(ctx *ThereforeContext) {}

// ExitTherefore is called when production Therefore is exited.
func (s *BaseFoxySheepListener) ExitTherefore(ctx *ThereforeContext) {}

// EnterTagUnset is called when production TagUnset is entered.
func (s *BaseFoxySheepListener) EnterTagUnset(ctx *TagUnsetContext) {}

// ExitTagUnset is called when production TagUnset is exited.
func (s *BaseFoxySheepListener) ExitTagUnset(ctx *TagUnsetContext) {}

// EnterAccessor is called when production Accessor is entered.
func (s *BaseFoxySheepListener) EnterAccessor(ctx *AccessorContext) {}

// ExitAccessor is called when production Accessor is exited.
func (s *BaseFoxySheepListener) ExitAccessor(ctx *AccessorContext) {}

// EnterCircleMinus is called when production CircleMinus is entered.
func (s *BaseFoxySheepListener) EnterCircleMinus(ctx *CircleMinusContext) {}

// ExitCircleMinus is called when production CircleMinus is exited.
func (s *BaseFoxySheepListener) ExitCircleMinus(ctx *CircleMinusContext) {}

// EnterDivide is called when production Divide is entered.
func (s *BaseFoxySheepListener) EnterDivide(ctx *DivideContext) {}

// ExitDivide is called when production Divide is exited.
func (s *BaseFoxySheepListener) ExitDivide(ctx *DivideContext) {}

// EnterImplies is called when production Implies is entered.
func (s *BaseFoxySheepListener) EnterImplies(ctx *ImpliesContext) {}

// ExitImplies is called when production Implies is exited.
func (s *BaseFoxySheepListener) ExitImplies(ctx *ImpliesContext) {}

// EnterPlusOp is called when production PlusOp is entered.
func (s *BaseFoxySheepListener) EnterPlusOp(ctx *PlusOpContext) {}

// ExitPlusOp is called when production PlusOp is exited.
func (s *BaseFoxySheepListener) ExitPlusOp(ctx *PlusOpContext) {}

// EnterRightComposition is called when production RightComposition is entered.
func (s *BaseFoxySheepListener) EnterRightComposition(ctx *RightCompositionContext) {}

// ExitRightComposition is called when production RightComposition is exited.
func (s *BaseFoxySheepListener) ExitRightComposition(ctx *RightCompositionContext) {}

// EnterNonCommutativeMultiply is called when production NonCommutativeMultiply is entered.
func (s *BaseFoxySheepListener) EnterNonCommutativeMultiply(ctx *NonCommutativeMultiplyContext) {}

// ExitNonCommutativeMultiply is called when production NonCommutativeMultiply is exited.
func (s *BaseFoxySheepListener) ExitNonCommutativeMultiply(ctx *NonCommutativeMultiplyContext) {}

// EnterList is called when production List is entered.
func (s *BaseFoxySheepListener) EnterList(ctx *ListContext) {}

// ExitList is called when production List is exited.
func (s *BaseFoxySheepListener) ExitList(ctx *ListContext) {}

// EnterCup is called when production Cup is entered.
func (s *BaseFoxySheepListener) EnterCup(ctx *CupContext) {}

// ExitCup is called when production Cup is exited.
func (s *BaseFoxySheepListener) ExitCup(ctx *CupContext) {}

// EnterSame is called when production Same is entered.
func (s *BaseFoxySheepListener) EnterSame(ctx *SameContext) {}

// ExitSame is called when production Same is exited.
func (s *BaseFoxySheepListener) ExitSame(ctx *SameContext) {}

// EnterBoxParen is called when production BoxParen is entered.
func (s *BaseFoxySheepListener) EnterBoxParen(ctx *BoxParenContext) {}

// ExitBoxParen is called when production BoxParen is exited.
func (s *BaseFoxySheepListener) ExitBoxParen(ctx *BoxParenContext) {}

// EnterOptional is called when production Optional is entered.
func (s *BaseFoxySheepListener) EnterOptional(ctx *OptionalContext) {}

// ExitOptional is called when production Optional is exited.
func (s *BaseFoxySheepListener) ExitOptional(ctx *OptionalContext) {}

// EnterSuchThat is called when production SuchThat is entered.
func (s *BaseFoxySheepListener) EnterSuchThat(ctx *SuchThatContext) {}

// ExitSuchThat is called when production SuchThat is exited.
func (s *BaseFoxySheepListener) ExitSuchThat(ctx *SuchThatContext) {}

// EnterDoubleBracketingBar is called when production DoubleBracketingBar is entered.
func (s *BaseFoxySheepListener) EnterDoubleBracketingBar(ctx *DoubleBracketingBarContext) {}

// ExitDoubleBracketingBar is called when production DoubleBracketingBar is exited.
func (s *BaseFoxySheepListener) ExitDoubleBracketingBar(ctx *DoubleBracketingBarContext) {}

// EnterPatternBlankDot is called when production PatternBlankDot is entered.
func (s *BaseFoxySheepListener) EnterPatternBlankDot(ctx *PatternBlankDotContext) {}

// ExitPatternBlankDot is called when production PatternBlankDot is exited.
func (s *BaseFoxySheepListener) ExitPatternBlankDot(ctx *PatternBlankDotContext) {}

// EnterDot is called when production Dot is entered.
func (s *BaseFoxySheepListener) EnterDot(ctx *DotContext) {}

// ExitDot is called when production Dot is exited.
func (s *BaseFoxySheepListener) ExitDot(ctx *DotContext) {}

// EnterVerticalBar is called when production VerticalBar is entered.
func (s *BaseFoxySheepListener) EnterVerticalBar(ctx *VerticalBarContext) {}

// ExitVerticalBar is called when production VerticalBar is exited.
func (s *BaseFoxySheepListener) ExitVerticalBar(ctx *VerticalBarContext) {}

// EnterSquare is called when production Square is entered.
func (s *BaseFoxySheepListener) EnterSquare(ctx *SquareContext) {}

// ExitSquare is called when production Square is exited.
func (s *BaseFoxySheepListener) ExitSquare(ctx *SquareContext) {}

// EnterAlternatives is called when production Alternatives is entered.
func (s *BaseFoxySheepListener) EnterAlternatives(ctx *AlternativesContext) {}

// ExitAlternatives is called when production Alternatives is exited.
func (s *BaseFoxySheepListener) ExitAlternatives(ctx *AlternativesContext) {}

// EnterOut is called when production Out is entered.
func (s *BaseFoxySheepListener) EnterOut(ctx *OutContext) {}

// ExitOut is called when production Out is exited.
func (s *BaseFoxySheepListener) ExitOut(ctx *OutContext) {}

// EnterBoxConstructor is called when production BoxConstructor is entered.
func (s *BaseFoxySheepListener) EnterBoxConstructor(ctx *BoxConstructorContext) {}

// ExitBoxConstructor is called when production BoxConstructor is exited.
func (s *BaseFoxySheepListener) ExitBoxConstructor(ctx *BoxConstructorContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseFoxySheepListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseFoxySheepListener) ExitNot(ctx *NotContext) {}

// EnterPostfix is called when production Postfix is entered.
func (s *BaseFoxySheepListener) EnterPostfix(ctx *PostfixContext) {}

// ExitPostfix is called when production Postfix is exited.
func (s *BaseFoxySheepListener) ExitPostfix(ctx *PostfixContext) {}

// EnterPatternBlanks is called when production PatternBlanks is entered.
func (s *BaseFoxySheepListener) EnterPatternBlanks(ctx *PatternBlanksContext) {}

// ExitPatternBlanks is called when production PatternBlanks is exited.
func (s *BaseFoxySheepListener) ExitPatternBlanks(ctx *PatternBlanksContext) {}

// EnterUnaryPlusMinus is called when production UnaryPlusMinus is entered.
func (s *BaseFoxySheepListener) EnterUnaryPlusMinus(ctx *UnaryPlusMinusContext) {}

// ExitUnaryPlusMinus is called when production UnaryPlusMinus is exited.
func (s *BaseFoxySheepListener) ExitUnaryPlusMinus(ctx *UnaryPlusMinusContext) {}

// EnterCap is called when production Cap is entered.
func (s *BaseFoxySheepListener) EnterCap(ctx *CapContext) {}

// ExitCap is called when production Cap is exited.
func (s *BaseFoxySheepListener) ExitCap(ctx *CapContext) {}

// EnterCirclePlus is called when production CirclePlus is entered.
func (s *BaseFoxySheepListener) EnterCirclePlus(ctx *CirclePlusContext) {}

// ExitCirclePlus is called when production CirclePlus is exited.
func (s *BaseFoxySheepListener) ExitCirclePlus(ctx *CirclePlusContext) {}

// EnterBecause is called when production Because is entered.
func (s *BaseFoxySheepListener) EnterBecause(ctx *BecauseContext) {}

// ExitBecause is called when production Because is exited.
func (s *BaseFoxySheepListener) ExitBecause(ctx *BecauseContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseFoxySheepListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseFoxySheepListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseFoxySheepListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseFoxySheepListener) ExitAnd(ctx *AndContext) {}

// EnterGet is called when production Get is entered.
func (s *BaseFoxySheepListener) EnterGet(ctx *GetContext) {}

// ExitGet is called when production Get is exited.
func (s *BaseFoxySheepListener) ExitGet(ctx *GetContext) {}

// EnterEquivalent is called when production Equivalent is entered.
func (s *BaseFoxySheepListener) EnterEquivalent(ctx *EquivalentContext) {}

// ExitEquivalent is called when production Equivalent is exited.
func (s *BaseFoxySheepListener) ExitEquivalent(ctx *EquivalentContext) {}

// EnterCompoundExpression is called when production CompoundExpression is entered.
func (s *BaseFoxySheepListener) EnterCompoundExpression(ctx *CompoundExpressionContext) {}

// ExitCompoundExpression is called when production CompoundExpression is exited.
func (s *BaseFoxySheepListener) ExitCompoundExpression(ctx *CompoundExpressionContext) {}

// EnterDerivative is called when production Derivative is entered.
func (s *BaseFoxySheepListener) EnterDerivative(ctx *DerivativeContext) {}

// ExitDerivative is called when production Derivative is exited.
func (s *BaseFoxySheepListener) ExitDerivative(ctx *DerivativeContext) {}

// EnterSlot is called when production Slot is entered.
func (s *BaseFoxySheepListener) EnterSlot(ctx *SlotContext) {}

// ExitSlot is called when production Slot is exited.
func (s *BaseFoxySheepListener) ExitSlot(ctx *SlotContext) {}

// EnterRightTee is called when production RightTee is entered.
func (s *BaseFoxySheepListener) EnterRightTee(ctx *RightTeeContext) {}

// ExitRightTee is called when production RightTee is exited.
func (s *BaseFoxySheepListener) ExitRightTee(ctx *RightTeeContext) {}

// EnterXor is called when production Xor is entered.
func (s *BaseFoxySheepListener) EnterXor(ctx *XorContext) {}

// ExitXor is called when production Xor is exited.
func (s *BaseFoxySheepListener) ExitXor(ctx *XorContext) {}

// EnterRule is called when production Rule is entered.
func (s *BaseFoxySheepListener) EnterRule(ctx *RuleContext) {}

// ExitRule is called when production Rule is exited.
func (s *BaseFoxySheepListener) ExitRule(ctx *RuleContext) {}

// EnterHeadExpression is called when production HeadExpression is entered.
func (s *BaseFoxySheepListener) EnterHeadExpression(ctx *HeadExpressionContext) {}

// ExitHeadExpression is called when production HeadExpression is exited.
func (s *BaseFoxySheepListener) ExitHeadExpression(ctx *HeadExpressionContext) {}

// EnterReplaceAll is called when production ReplaceAll is entered.
func (s *BaseFoxySheepListener) EnterReplaceAll(ctx *ReplaceAllContext) {}

// ExitReplaceAll is called when production ReplaceAll is exited.
func (s *BaseFoxySheepListener) ExitReplaceAll(ctx *ReplaceAllContext) {}

// EnterIntersection is called when production Intersection is entered.
func (s *BaseFoxySheepListener) EnterIntersection(ctx *IntersectionContext) {}

// ExitIntersection is called when production Intersection is exited.
func (s *BaseFoxySheepListener) ExitIntersection(ctx *IntersectionContext) {}

// EnterPreIncrement is called when production PreIncrement is entered.
func (s *BaseFoxySheepListener) EnterPreIncrement(ctx *PreIncrementContext) {}

// ExitPreIncrement is called when production PreIncrement is exited.
func (s *BaseFoxySheepListener) ExitPreIncrement(ctx *PreIncrementContext) {}

// EnterIntegrate is called when production Integrate is entered.
func (s *BaseFoxySheepListener) EnterIntegrate(ctx *IntegrateContext) {}

// ExitIntegrate is called when production Integrate is exited.
func (s *BaseFoxySheepListener) ExitIntegrate(ctx *IntegrateContext) {}

// EnterSet is called when production Set is entered.
func (s *BaseFoxySheepListener) EnterSet(ctx *SetContext) {}

// ExitSet is called when production Set is exited.
func (s *BaseFoxySheepListener) ExitSet(ctx *SetContext) {}

// EnterOpEquals is called when production OpEquals is entered.
func (s *BaseFoxySheepListener) EnterOpEquals(ctx *OpEqualsContext) {}

// ExitOpEquals is called when production OpEquals is exited.
func (s *BaseFoxySheepListener) ExitOpEquals(ctx *OpEqualsContext) {}

// EnterMessage is called when production Message is entered.
func (s *BaseFoxySheepListener) EnterMessage(ctx *MessageContext) {}

// ExitMessage is called when production Message is exited.
func (s *BaseFoxySheepListener) ExitMessage(ctx *MessageContext) {}

// EnterCross is called when production Cross is entered.
func (s *BaseFoxySheepListener) EnterCross(ctx *CrossContext) {}

// ExitCross is called when production Cross is exited.
func (s *BaseFoxySheepListener) ExitCross(ctx *CrossContext) {}

// EnterPatternTest is called when production PatternTest is entered.
func (s *BaseFoxySheepListener) EnterPatternTest(ctx *PatternTestContext) {}

// ExitPatternTest is called when production PatternTest is exited.
func (s *BaseFoxySheepListener) ExitPatternTest(ctx *PatternTestContext) {}

// EnterPrefix is called when production Prefix is entered.
func (s *BaseFoxySheepListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production Prefix is exited.
func (s *BaseFoxySheepListener) ExitPrefix(ctx *PrefixContext) {}

// EnterBackslash is called when production Backslash is entered.
func (s *BaseFoxySheepListener) EnterBackslash(ctx *BackslashContext) {}

// ExitBackslash is called when production Backslash is exited.
func (s *BaseFoxySheepListener) ExitBackslash(ctx *BackslashContext) {}

// EnterRepeated is called when production Repeated is entered.
func (s *BaseFoxySheepListener) EnterRepeated(ctx *RepeatedContext) {}

// ExitRepeated is called when production Repeated is exited.
func (s *BaseFoxySheepListener) ExitRepeated(ctx *RepeatedContext) {}

// EnterMapApply is called when production MapApply is entered.
func (s *BaseFoxySheepListener) EnterMapApply(ctx *MapApplyContext) {}

// ExitMapApply is called when production MapApply is exited.
func (s *BaseFoxySheepListener) ExitMapApply(ctx *MapApplyContext) {}

// EnterUnion is called when production Union is entered.
func (s *BaseFoxySheepListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production Union is exited.
func (s *BaseFoxySheepListener) ExitUnion(ctx *UnionContext) {}

// EnterVerticalSeparator is called when production VerticalSeparator is entered.
func (s *BaseFoxySheepListener) EnterVerticalSeparator(ctx *VerticalSeparatorContext) {}

// ExitVerticalSeparator is called when production VerticalSeparator is exited.
func (s *BaseFoxySheepListener) ExitVerticalSeparator(ctx *VerticalSeparatorContext) {}

// EnterFactorial is called when production Factorial is entered.
func (s *BaseFoxySheepListener) EnterFactorial(ctx *FactorialContext) {}

// ExitFactorial is called when production Factorial is exited.
func (s *BaseFoxySheepListener) ExitFactorial(ctx *FactorialContext) {}

// EnterSpanA is called when production SpanA is entered.
func (s *BaseFoxySheepListener) EnterSpanA(ctx *SpanAContext) {}

// ExitSpanA is called when production SpanA is exited.
func (s *BaseFoxySheepListener) ExitSpanA(ctx *SpanAContext) {}

// EnterFunction is called when production Function is entered.
func (s *BaseFoxySheepListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production Function is exited.
func (s *BaseFoxySheepListener) ExitFunction(ctx *FunctionContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseFoxySheepListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseFoxySheepListener) ExitNumber(ctx *NumberContext) {}

// EnterStar is called when production Star is entered.
func (s *BaseFoxySheepListener) EnterStar(ctx *StarContext) {}

// ExitStar is called when production Star is exited.
func (s *BaseFoxySheepListener) ExitStar(ctx *StarContext) {}

// EnterComparison is called when production Comparison is entered.
func (s *BaseFoxySheepListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production Comparison is exited.
func (s *BaseFoxySheepListener) ExitComparison(ctx *ComparisonContext) {}

// EnterTagSet is called when production TagSet is entered.
func (s *BaseFoxySheepListener) EnterTagSet(ctx *TagSetContext) {}

// ExitTagSet is called when production TagSet is exited.
func (s *BaseFoxySheepListener) ExitTagSet(ctx *TagSetContext) {}

// EnterIncrement is called when production Increment is entered.
func (s *BaseFoxySheepListener) EnterIncrement(ctx *IncrementContext) {}

// ExitIncrement is called when production Increment is exited.
func (s *BaseFoxySheepListener) ExitIncrement(ctx *IncrementContext) {}

// EnterVerticalTilde is called when production VerticalTilde is entered.
func (s *BaseFoxySheepListener) EnterVerticalTilde(ctx *VerticalTildeContext) {}

// ExitVerticalTilde is called when production VerticalTilde is exited.
func (s *BaseFoxySheepListener) ExitVerticalTilde(ctx *VerticalTildeContext) {}

// EnterColon is called when production Colon is entered.
func (s *BaseFoxySheepListener) EnterColon(ctx *ColonContext) {}

// ExitColon is called when production Colon is exited.
func (s *BaseFoxySheepListener) ExitColon(ctx *ColonContext) {}

// EnterSmallCircle is called when production SmallCircle is entered.
func (s *BaseFoxySheepListener) EnterSmallCircle(ctx *SmallCircleContext) {}

// ExitSmallCircle is called when production SmallCircle is exited.
func (s *BaseFoxySheepListener) ExitSmallCircle(ctx *SmallCircleContext) {}

// EnterParentheses is called when production Parentheses is entered.
func (s *BaseFoxySheepListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production Parentheses is exited.
func (s *BaseFoxySheepListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterSpanB is called when production SpanB is entered.
func (s *BaseFoxySheepListener) EnterSpanB(ctx *SpanBContext) {}

// ExitSpanB is called when production SpanB is exited.
func (s *BaseFoxySheepListener) ExitSpanB(ctx *SpanBContext) {}

// EnterCondition is called when production Condition is entered.
func (s *BaseFoxySheepListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production Condition is exited.
func (s *BaseFoxySheepListener) ExitCondition(ctx *ConditionContext) {}

// EnterFloor is called when production Floor is entered.
func (s *BaseFoxySheepListener) EnterFloor(ctx *FloorContext) {}

// ExitFloor is called when production Floor is exited.
func (s *BaseFoxySheepListener) ExitFloor(ctx *FloorContext) {}

// EnterComposition is called when production Composition is entered.
func (s *BaseFoxySheepListener) EnterComposition(ctx *CompositionContext) {}

// ExitComposition is called when production Composition is exited.
func (s *BaseFoxySheepListener) ExitComposition(ctx *CompositionContext) {}

// EnterCircleDot is called when production CircleDot is entered.
func (s *BaseFoxySheepListener) EnterCircleDot(ctx *CircleDotContext) {}

// ExitCircleDot is called when production CircleDot is exited.
func (s *BaseFoxySheepListener) ExitCircleDot(ctx *CircleDotContext) {}

// EnterSymbolLiteral is called when production SymbolLiteral is entered.
func (s *BaseFoxySheepListener) EnterSymbolLiteral(ctx *SymbolLiteralContext) {}

// ExitSymbolLiteral is called when production SymbolLiteral is exited.
func (s *BaseFoxySheepListener) ExitSymbolLiteral(ctx *SymbolLiteralContext) {}

// EnterCircleTimes is called when production CircleTimes is entered.
func (s *BaseFoxySheepListener) EnterCircleTimes(ctx *CircleTimesContext) {}

// ExitCircleTimes is called when production CircleTimes is exited.
func (s *BaseFoxySheepListener) ExitCircleTimes(ctx *CircleTimesContext) {}

// EnterUnset is called when production Unset is entered.
func (s *BaseFoxySheepListener) EnterUnset(ctx *UnsetContext) {}

// ExitUnset is called when production Unset is exited.
func (s *BaseFoxySheepListener) ExitUnset(ctx *UnsetContext) {}

// EnterDel is called when production Del is entered.
func (s *BaseFoxySheepListener) EnterDel(ctx *DelContext) {}

// ExitDel is called when production Del is exited.
func (s *BaseFoxySheepListener) ExitDel(ctx *DelContext) {}

// EnterDiamond is called when production Diamond is entered.
func (s *BaseFoxySheepListener) EnterDiamond(ctx *DiamondContext) {}

// ExitDiamond is called when production Diamond is exited.
func (s *BaseFoxySheepListener) ExitDiamond(ctx *DiamondContext) {}

// EnterWedge is called when production Wedge is entered.
func (s *BaseFoxySheepListener) EnterWedge(ctx *WedgeContext) {}

// ExitWedge is called when production Wedge is exited.
func (s *BaseFoxySheepListener) ExitWedge(ctx *WedgeContext) {}

// EnterPut is called when production Put is entered.
func (s *BaseFoxySheepListener) EnterPut(ctx *PutContext) {}

// ExitPut is called when production Put is exited.
func (s *BaseFoxySheepListener) ExitPut(ctx *PutContext) {}

// EnterStringJoin is called when production StringJoin is entered.
func (s *BaseFoxySheepListener) EnterStringJoin(ctx *StringJoinContext) {}

// ExitStringJoin is called when production StringJoin is exited.
func (s *BaseFoxySheepListener) ExitStringJoin(ctx *StringJoinContext) {}

// EnterTee is called when production Tee is entered.
func (s *BaseFoxySheepListener) EnterTee(ctx *TeeContext) {}

// ExitTee is called when production Tee is exited.
func (s *BaseFoxySheepListener) ExitTee(ctx *TeeContext) {}

// EnterSetContainment is called when production SetContainment is entered.
func (s *BaseFoxySheepListener) EnterSetContainment(ctx *SetContainmentContext) {}

// ExitSetContainment is called when production SetContainment is exited.
func (s *BaseFoxySheepListener) ExitSetContainment(ctx *SetContainmentContext) {}

// EnterVee is called when production Vee is entered.
func (s *BaseFoxySheepListener) EnterVee(ctx *VeeContext) {}

// ExitVee is called when production Vee is exited.
func (s *BaseFoxySheepListener) ExitVee(ctx *VeeContext) {}

// EnterCenterDot is called when production CenterDot is entered.
func (s *BaseFoxySheepListener) EnterCenterDot(ctx *CenterDotContext) {}

// ExitCenterDot is called when production CenterDot is exited.
func (s *BaseFoxySheepListener) ExitCenterDot(ctx *CenterDotContext) {}

// EnterTimes is called when production Times is entered.
func (s *BaseFoxySheepListener) EnterTimes(ctx *TimesContext) {}

// ExitTimes is called when production Times is exited.
func (s *BaseFoxySheepListener) ExitTimes(ctx *TimesContext) {}

// EnterStringExpression is called when production StringExpression is entered.
func (s *BaseFoxySheepListener) EnterStringExpression(ctx *StringExpressionContext) {}

// ExitStringExpression is called when production StringExpression is exited.
func (s *BaseFoxySheepListener) ExitStringExpression(ctx *StringExpressionContext) {}

// EnterBracketingBar is called when production BracketingBar is entered.
func (s *BaseFoxySheepListener) EnterBracketingBar(ctx *BracketingBarContext) {}

// ExitBracketingBar is called when production BracketingBar is exited.
func (s *BaseFoxySheepListener) ExitBracketingBar(ctx *BracketingBarContext) {}

// EnterCoproduct is called when production Coproduct is entered.
func (s *BaseFoxySheepListener) EnterCoproduct(ctx *CoproductContext) {}

// ExitCoproduct is called when production Coproduct is exited.
func (s *BaseFoxySheepListener) ExitCoproduct(ctx *CoproductContext) {}

// EnterAngleBracket is called when production AngleBracket is entered.
func (s *BaseFoxySheepListener) EnterAngleBracket(ctx *AngleBracketContext) {}

// ExitAngleBracket is called when production AngleBracket is exited.
func (s *BaseFoxySheepListener) ExitAngleBracket(ctx *AngleBracketContext) {}

// EnterPower is called when production Power is entered.
func (s *BaseFoxySheepListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production Power is exited.
func (s *BaseFoxySheepListener) ExitPower(ctx *PowerContext) {}

// EnterContextName is called when production ContextName is entered.
func (s *BaseFoxySheepListener) EnterContextName(ctx *ContextNameContext) {}

// ExitContextName is called when production ContextName is exited.
func (s *BaseFoxySheepListener) ExitContextName(ctx *ContextNameContext) {}

// EnterSimpleContext is called when production SimpleContext is entered.
func (s *BaseFoxySheepListener) EnterSimpleContext(ctx *SimpleContextContext) {}

// ExitSimpleContext is called when production SimpleContext is exited.
func (s *BaseFoxySheepListener) ExitSimpleContext(ctx *SimpleContextContext) {}

// EnterCompoundContext is called when production CompoundContext is entered.
func (s *BaseFoxySheepListener) EnterCompoundContext(ctx *CompoundContextContext) {}

// ExitCompoundContext is called when production CompoundContext is exited.
func (s *BaseFoxySheepListener) ExitCompoundContext(ctx *CompoundContextContext) {}

// EnterNumberBaseN is called when production NumberBaseN is entered.
func (s *BaseFoxySheepListener) EnterNumberBaseN(ctx *NumberBaseNContext) {}

// ExitNumberBaseN is called when production NumberBaseN is exited.
func (s *BaseFoxySheepListener) ExitNumberBaseN(ctx *NumberBaseNContext) {}

// EnterNumberBaseTen is called when production NumberBaseTen is entered.
func (s *BaseFoxySheepListener) EnterNumberBaseTen(ctx *NumberBaseTenContext) {}

// ExitNumberBaseTen is called when production NumberBaseTen is exited.
func (s *BaseFoxySheepListener) ExitNumberBaseTen(ctx *NumberBaseTenContext) {}

// EnterNumberLiteralPrecision is called when production numberLiteralPrecision is entered.
func (s *BaseFoxySheepListener) EnterNumberLiteralPrecision(ctx *NumberLiteralPrecisionContext) {}

// ExitNumberLiteralPrecision is called when production numberLiteralPrecision is exited.
func (s *BaseFoxySheepListener) ExitNumberLiteralPrecision(ctx *NumberLiteralPrecisionContext) {}

// EnterNumberLiteralExponent is called when production numberLiteralExponent is entered.
func (s *BaseFoxySheepListener) EnterNumberLiteralExponent(ctx *NumberLiteralExponentContext) {}

// ExitNumberLiteralExponent is called when production numberLiteralExponent is exited.
func (s *BaseFoxySheepListener) ExitNumberLiteralExponent(ctx *NumberLiteralExponentContext) {}

// EnterOutNumbered is called when production OutNumbered is entered.
func (s *BaseFoxySheepListener) EnterOutNumbered(ctx *OutNumberedContext) {}

// ExitOutNumbered is called when production OutNumbered is exited.
func (s *BaseFoxySheepListener) ExitOutNumbered(ctx *OutNumberedContext) {}

// EnterOutUnnumbered is called when production OutUnnumbered is entered.
func (s *BaseFoxySheepListener) EnterOutUnnumbered(ctx *OutUnnumberedContext) {}

// ExitOutUnnumbered is called when production OutUnnumbered is exited.
func (s *BaseFoxySheepListener) ExitOutUnnumbered(ctx *OutUnnumberedContext) {}

// EnterSlotDigits is called when production SlotDigits is entered.
func (s *BaseFoxySheepListener) EnterSlotDigits(ctx *SlotDigitsContext) {}

// ExitSlotDigits is called when production SlotDigits is exited.
func (s *BaseFoxySheepListener) ExitSlotDigits(ctx *SlotDigitsContext) {}

// EnterSlotNamed is called when production SlotNamed is entered.
func (s *BaseFoxySheepListener) EnterSlotNamed(ctx *SlotNamedContext) {}

// ExitSlotNamed is called when production SlotNamed is exited.
func (s *BaseFoxySheepListener) ExitSlotNamed(ctx *SlotNamedContext) {}

// EnterSlotSequenceDigits is called when production SlotSequenceDigits is entered.
func (s *BaseFoxySheepListener) EnterSlotSequenceDigits(ctx *SlotSequenceDigitsContext) {}

// ExitSlotSequenceDigits is called when production SlotSequenceDigits is exited.
func (s *BaseFoxySheepListener) ExitSlotSequenceDigits(ctx *SlotSequenceDigitsContext) {}

// EnterSlotSequence is called when production SlotSequence is entered.
func (s *BaseFoxySheepListener) EnterSlotSequence(ctx *SlotSequenceContext) {}

// ExitSlotSequence is called when production SlotSequence is exited.
func (s *BaseFoxySheepListener) ExitSlotSequence(ctx *SlotSequenceContext) {}

// EnterSlotUnnamed is called when production SlotUnnamed is entered.
func (s *BaseFoxySheepListener) EnterSlotUnnamed(ctx *SlotUnnamedContext) {}

// ExitSlotUnnamed is called when production SlotUnnamed is exited.
func (s *BaseFoxySheepListener) ExitSlotUnnamed(ctx *SlotUnnamedContext) {}

// EnterExpressionListed is called when production ExpressionListed is entered.
func (s *BaseFoxySheepListener) EnterExpressionListed(ctx *ExpressionListedContext) {}

// ExitExpressionListed is called when production ExpressionListed is exited.
func (s *BaseFoxySheepListener) ExitExpressionListed(ctx *ExpressionListedContext) {}

// EnterAccessExpressionA is called when production AccessExpressionA is entered.
func (s *BaseFoxySheepListener) EnterAccessExpressionA(ctx *AccessExpressionAContext) {}

// ExitAccessExpressionA is called when production AccessExpressionA is exited.
func (s *BaseFoxySheepListener) ExitAccessExpressionA(ctx *AccessExpressionAContext) {}

// EnterAccessExpressionB is called when production AccessExpressionB is entered.
func (s *BaseFoxySheepListener) EnterAccessExpressionB(ctx *AccessExpressionBContext) {}

// ExitAccessExpressionB is called when production AccessExpressionB is exited.
func (s *BaseFoxySheepListener) ExitAccessExpressionB(ctx *AccessExpressionBContext) {}

// EnterBox is called when production box is entered.
func (s *BaseFoxySheepListener) EnterBox(ctx *BoxContext) {}

// ExitBox is called when production box is exited.
func (s *BaseFoxySheepListener) ExitBox(ctx *BoxContext) {}
