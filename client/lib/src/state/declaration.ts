import { IsEnum, IsInt, IsOptional, ValidateNested } from 'class-validator'
import { Type } from 'class-transformer'
import { Card } from '../card'

export enum DeclarationType {
	Tierce = 'Tierce',
	Quarte = 'Quarte',
	Quinte = 'Quinte',
	Carre = 'Carre',
	NinesCarre = 'NinesCarre',
	JacksCarre = 'JacksCarre',
	Belote = 'Belote',
}

export function getDeclarationTypeName(type: DeclarationType): string {
	switch (type) {
		case DeclarationType.Tierce:
			return 'Tierce'
		case DeclarationType.Quarte:
			return 'Quarte'
		case DeclarationType.Quinte:
			return 'Quinte'
		case DeclarationType.Carre:
			return 'Carré'
		case DeclarationType.NinesCarre:
			return 'Carré'
		case DeclarationType.JacksCarre:
			return 'Carré'
		case DeclarationType.Belote:
			return 'Belote'
	}
}

export class Declaration {
	@IsEnum(DeclarationType)
	type: DeclarationType

	@Type(() => Card)
	@IsOptional()
	@ValidateNested()
	highestCard?: Card

	@IsInt()
	points: number

	constructor(type: DeclarationType, points: number, highestCard?: Card) {
		this.type = type
		this.points = points
		this.highestCard = highestCard
	}
}
