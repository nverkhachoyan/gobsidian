---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/14
- monster/environment/desert
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githyanki Supreme Commander"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 206
---
# [Githyanki Supreme Commander](3-Mechanics\CLI\bestiary\humanoid/githyanki-supreme-commander-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 206*  

> [!quote]- A quote from Mordenkainen  
> 
> What would become of this multiverse if githyanki didn't guard the Astral Plane from the illithid menace? What would reality become if beings of thought ruled the plane of thought?

## Githyanki Supreme Commander

Supreme commanders lead the githyanki armies, each one commanding ten kith'raks, who in turn lead the rest of their forces. Most supreme commanders ride red dragons into battle.

## Gith

The descendants of an ancient people—so old their original name has been lost—have turned against each other, becoming vicious enemies divided over mortality, purpose, and the machinations of their leaders. The bellicose githyanki terrorize the Astral Plane, raiding into other worlds to plunder the multiverse of its magic and riches. The githzerai live apart from the rest of the cosmos, content within the confines of their fortresses floating through the chaos of Limbo. Although the two groups of gith despise each other, their hatred for the mind flayers from whom they escaped endures, and both githyanki and githzerai are dedicated to hunting their ancestral foes.

```statblock
"name": "Githyanki Supreme Commander (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Lawful Evil"
"ac": !!int "18"
"ac_class": "[plate armor](/3-Mechanics/CLI/items/plate-armor.md)"
"hp": !!int "187"
"hit_dice": "22d8 + 88"
"stats":
- !!int "19"
- !!int "17"
- !!int "18"
- !!int "16"
- !!int "16"
- !!int "18"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "8"
  "Intelligence": !!int "8"
  "Constitution": !!int "9"
"skillsaves":
  "Intimidation": !!int "9"
  "Insight": !!int "8"
  "Perception": !!int "8"
"senses": "passive Perception 18"
"languages": "Gith"
"cr": "14"
"traits":
- "desc": "The githyanki's innate spellcasting ability is Intelligence (spell save\
    \ DC 16, dice: d20+8 (+8) to hit with spell attacks). It can innately cast\
    \ the following spells, requiring no components:\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\
    \ (the hand is invisible)\n\n1/day each: [Bigby's hand](/3-Mechanics/CLI/spells/bigbys-hand.md),\
    \ [mass suggestion](/3-Mechanics/CLI/spells/mass-suggestion.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\n3/day each: [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md) (self only), [misty step](/3-Mechanics/CLI/spells/misty-step.md),\
    \ [nondetection](/3-Mechanics/CLI/spells/nondetection.md) (self only)"
  "name": "Innate Spellcasting (Psionics)"
"actions":
- "desc": "The githyanki makes two greatsword attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 7|text(14) (2d6 + 7) slashing damage plus dice:5d6|text(17)\
    \ (5d6) psychic damage. On a critical hit against a target in an astral body\
    \ (as with the [astral projection](/3-Mechanics/CLI/spells/astral-projection.md)\
    \ spell), the githyanki can cut the silvery cord that tethers the target to its\
    \ material body, instead of dealing damage."
  "name": "Silver Greatsword"
"reactions":
- "desc": "The githyanki adds 5 to its AC against one melee attack that would hit\
    \ it. To do so, the githyanki must see the attacker and be wielding a melee weapon."
  "name": "Parry"
"legendary_actions":
- "desc": "The githyanki makes a greatsword attack."
  "name": "Attack (2 Actions)"
- "desc": "The githyanki targets one ally it can see within 30 feet of it. If the\
    \ target can see or hear the githyanki, the target can make one melee weapon attack\
    \ using its reaction and has advantage on the attack roll."
  "name": "Command Ally"
- "desc": "The githyanki magically teleports, along with any equipment it is wearing\
    \ and carrying, to an unoccupied space it can see within 30 feet of it. It also\
    \ becomes insubstantial until the start of its next turn. While insubstantial,\
    \ it can move through other creatures and objects as if they were difficult terrain.\
    \ If it ends its turn inside an object, it takes dice:3d10|text(16) (3d10)\
    \ force damage and is moved to the nearest unoccupied space."
  "name": "Teleport"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Githyanki%20Supreme%20Commander.webp"
```
^statblock

```encounter-table
name: Githyanki Supreme Commander
creatures:
 - 1: Githyanki Supreme Commander
```

## Environment

desert, mountain, urban