---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/12
- monster/environment/desert
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githyanki Kith'rak"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 205
---
# [Githyanki Kith'rak](3-Mechanics\CLI\bestiary\humanoid/githyanki-kithrak-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 205*  

> [!quote]- A quote from Mordenkainen  
> 
> What would become of this multiverse if githyanki didn't guard the Astral Plane from the illithid menace? What would reality become if beings of thought ruled the plane of thought?

## Githyanki Kith'rak

The githyanki's militarized culture assigns ranks and responsibilities to its citizens. Groups of ten warriors follow the commands of the sarths (githyanki warriors), while ten sarths obey the commands of the mighty kith'rak. These champions earn their status through torturous training and psionic testing until they can command the respect of their underlings.

## Gith

The descendants of an ancient people—so old their original name has been lost—have turned against each other, becoming vicious enemies divided over mortality, purpose, and the machinations of their leaders. The bellicose githyanki terrorize the Astral Plane, raiding into other worlds to plunder the multiverse of its magic and riches. The githzerai live apart from the rest of the cosmos, content within the confines of their fortresses floating through the chaos of Limbo. Although the two groups of gith despise each other, their hatred for the mind flayers from whom they escaped endures, and both githyanki and githzerai are dedicated to hunting their ancestral foes.

```statblock
"name": "Githyanki Kith'rak (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "180"
"hit_dice": "24d8 + 72"
"stats":
- !!int "18"
- !!int "16"
- !!int "17"
- !!int "16"
- !!int "15"
- !!int "17"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "6"
  "Intelligence": !!int "7"
  "Constitution": !!int "7"
"skillsaves":
  "Intimidation": !!int "7"
  "Perception": !!int "6"
"senses": "passive Perception 16"
"languages": "Gith"
"cr": "12"
"traits":
- "desc": "The githyanki's innate spellcasting ability is Intelligence (spell save\
    \ DC 15, dice: d20+7 (+7) to hit with spell attacks). It can innately cast\
    \ the following spells, requiring no components:\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\
    \ (the hand is invisible)\n\n1/day each: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\n3/day each: [blur](/3-Mechanics/CLI/spells/blur.md),\
    \ [jump](/3-Mechanics/CLI/spells/jump.md), [misty step](/3-Mechanics/CLI/spells/misty-step.md),\
    \ [nondetection](/3-Mechanics/CLI/spells/nondetection.md) (self only)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "As a bonus action, the githyanki can magically end the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ and [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) conditions\
    \ on itself and each creature of its choice that it can see within 30 feet of\
    \ it."
  "name": "Rally the Troops"
"actions":
- "desc": "The githyanki makes three greatsword attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) slashing damage plus dice:5d6|text(17)\
    \ (5d6) psychic damage."
  "name": "Greatsword"
"reactions":
- "desc": "The githyanki adds 4 to its AC against one melee attack that would hit\
    \ it. To do so, the githyanki must see the attacker and be wielding a melee weapon."
  "name": "Parry"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Githyanki%20Kith%27rak.webp"
```
^statblock

```encounter-table
name: Githyanki Kith'rak
creatures:
 - 1: Githyanki Kith'rak
```

## Environment

desert, mountain, urban