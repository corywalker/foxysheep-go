// Generated from ../FoxySheep.g4 by ANTLR 4.7.

package parser // FoxySheep

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FoxySheepParser.
type FoxySheepVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FoxySheepParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#PatternExp.
	VisitPatternExp(ctx *PatternExpContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Conjugate.
	VisitConjugate(ctx *ConjugateContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Ceiling.
	VisitCeiling(ctx *CeilingContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Infix.
	VisitInfix(ctx *InfixContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Therefore.
	VisitTherefore(ctx *ThereforeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#TagUnset.
	VisitTagUnset(ctx *TagUnsetContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Accessor.
	VisitAccessor(ctx *AccessorContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CircleMinus.
	VisitCircleMinus(ctx *CircleMinusContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Divide.
	VisitDivide(ctx *DivideContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Implies.
	VisitImplies(ctx *ImpliesContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#PlusOp.
	VisitPlusOp(ctx *PlusOpContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#RightComposition.
	VisitRightComposition(ctx *RightCompositionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#NonCommutativeMultiply.
	VisitNonCommutativeMultiply(ctx *NonCommutativeMultiplyContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#List.
	VisitList(ctx *ListContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Cup.
	VisitCup(ctx *CupContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Same.
	VisitSame(ctx *SameContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#BoxParen.
	VisitBoxParen(ctx *BoxParenContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Optional.
	VisitOptional(ctx *OptionalContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SuchThat.
	VisitSuchThat(ctx *SuchThatContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#DoubleBracketingBar.
	VisitDoubleBracketingBar(ctx *DoubleBracketingBarContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#PatternBlankDot.
	VisitPatternBlankDot(ctx *PatternBlankDotContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Dot.
	VisitDot(ctx *DotContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#VerticalBar.
	VisitVerticalBar(ctx *VerticalBarContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Square.
	VisitSquare(ctx *SquareContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Alternatives.
	VisitAlternatives(ctx *AlternativesContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Out.
	VisitOut(ctx *OutContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#BoxConstructor.
	VisitBoxConstructor(ctx *BoxConstructorContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Not.
	VisitNot(ctx *NotContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Postfix.
	VisitPostfix(ctx *PostfixContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#PatternBlanks.
	VisitPatternBlanks(ctx *PatternBlanksContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#UnaryPlusMinus.
	VisitUnaryPlusMinus(ctx *UnaryPlusMinusContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Cap.
	VisitCap(ctx *CapContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CirclePlus.
	VisitCirclePlus(ctx *CirclePlusContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Because.
	VisitBecause(ctx *BecauseContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#And.
	VisitAnd(ctx *AndContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Get.
	VisitGet(ctx *GetContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Equivalent.
	VisitEquivalent(ctx *EquivalentContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CompoundExpression.
	VisitCompoundExpression(ctx *CompoundExpressionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Derivative.
	VisitDerivative(ctx *DerivativeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Slot.
	VisitSlot(ctx *SlotContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#RightTee.
	VisitRightTee(ctx *RightTeeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Xor.
	VisitXor(ctx *XorContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Rule.
	VisitRule(ctx *RuleContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#HeadExpression.
	VisitHeadExpression(ctx *HeadExpressionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#ReplaceAll.
	VisitReplaceAll(ctx *ReplaceAllContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Intersection.
	VisitIntersection(ctx *IntersectionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#PreIncrement.
	VisitPreIncrement(ctx *PreIncrementContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Integrate.
	VisitIntegrate(ctx *IntegrateContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Set.
	VisitSet(ctx *SetContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#OpEquals.
	VisitOpEquals(ctx *OpEqualsContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Cross.
	VisitCross(ctx *CrossContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#PatternTest.
	VisitPatternTest(ctx *PatternTestContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Prefix.
	VisitPrefix(ctx *PrefixContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Backslash.
	VisitBackslash(ctx *BackslashContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Repeated.
	VisitRepeated(ctx *RepeatedContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#MapApply.
	VisitMapApply(ctx *MapApplyContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Union.
	VisitUnion(ctx *UnionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#VerticalSeparator.
	VisitVerticalSeparator(ctx *VerticalSeparatorContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Factorial.
	VisitFactorial(ctx *FactorialContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SpanA.
	VisitSpanA(ctx *SpanAContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Star.
	VisitStar(ctx *StarContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#TagSet.
	VisitTagSet(ctx *TagSetContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Increment.
	VisitIncrement(ctx *IncrementContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#VerticalTilde.
	VisitVerticalTilde(ctx *VerticalTildeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Colon.
	VisitColon(ctx *ColonContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SmallCircle.
	VisitSmallCircle(ctx *SmallCircleContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Parentheses.
	VisitParentheses(ctx *ParenthesesContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SpanB.
	VisitSpanB(ctx *SpanBContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Floor.
	VisitFloor(ctx *FloorContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Composition.
	VisitComposition(ctx *CompositionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CircleDot.
	VisitCircleDot(ctx *CircleDotContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SymbolLiteral.
	VisitSymbolLiteral(ctx *SymbolLiteralContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CircleTimes.
	VisitCircleTimes(ctx *CircleTimesContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Unset.
	VisitUnset(ctx *UnsetContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Del.
	VisitDel(ctx *DelContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Diamond.
	VisitDiamond(ctx *DiamondContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Wedge.
	VisitWedge(ctx *WedgeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Put.
	VisitPut(ctx *PutContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#StringJoin.
	VisitStringJoin(ctx *StringJoinContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Tee.
	VisitTee(ctx *TeeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SetContainment.
	VisitSetContainment(ctx *SetContainmentContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Vee.
	VisitVee(ctx *VeeContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CenterDot.
	VisitCenterDot(ctx *CenterDotContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Times.
	VisitTimes(ctx *TimesContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#StringExpression.
	VisitStringExpression(ctx *StringExpressionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#BracketingBar.
	VisitBracketingBar(ctx *BracketingBarContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Coproduct.
	VisitCoproduct(ctx *CoproductContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#AngleBracket.
	VisitAngleBracket(ctx *AngleBracketContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#Power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#ContextName.
	VisitContextName(ctx *ContextNameContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SimpleContext.
	VisitSimpleContext(ctx *SimpleContextContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#CompoundContext.
	VisitCompoundContext(ctx *CompoundContextContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#NumberBaseN.
	VisitNumberBaseN(ctx *NumberBaseNContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#NumberBaseTen.
	VisitNumberBaseTen(ctx *NumberBaseTenContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#numberLiteralPrecision.
	VisitNumberLiteralPrecision(ctx *NumberLiteralPrecisionContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#numberLiteralExponent.
	VisitNumberLiteralExponent(ctx *NumberLiteralExponentContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#OutNumbered.
	VisitOutNumbered(ctx *OutNumberedContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#OutUnnumbered.
	VisitOutUnnumbered(ctx *OutUnnumberedContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SlotDigits.
	VisitSlotDigits(ctx *SlotDigitsContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SlotNamed.
	VisitSlotNamed(ctx *SlotNamedContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SlotSequenceDigits.
	VisitSlotSequenceDigits(ctx *SlotSequenceDigitsContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SlotSequence.
	VisitSlotSequence(ctx *SlotSequenceContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#SlotUnnamed.
	VisitSlotUnnamed(ctx *SlotUnnamedContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#ExpressionListed.
	VisitExpressionListed(ctx *ExpressionListedContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#AccessExpressionA.
	VisitAccessExpressionA(ctx *AccessExpressionAContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#AccessExpressionB.
	VisitAccessExpressionB(ctx *AccessExpressionBContext) interface{}

	// Visit a parse tree produced by FoxySheepParser#box.
	VisitBox(ctx *BoxContext) interface{}
}
