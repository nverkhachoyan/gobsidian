---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/10
- monster/environment/desert
- monster/environment/mountain
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/gith
aliases: ["Githzerai Enlightened"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 208
---
# [Githzerai Enlightened](3-Mechanics\CLI\bestiary\humanoid/githzerai-enlightened-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 208*  

> [!quote]- A quote from Mordenkainen  
> 
> The githzerai are a check on the githyanki and the illithids. The githyanki are a check on the githzerai and the illithids. Thus, three unequal forces enforce the Balance.

## Githzerai Enlightened

Githzerai never stop training. They spend long hours in meditation to transcend the limits of their forms and to apprehend the nature of reality. Zerths who complete the next tier of their training become one of the githzerai known as the enlightened.

## Gith

The descendants of an ancient people—so old their original name has been lost—have turned against each other, becoming vicious enemies divided over mortality, purpose, and the machinations of their leaders. The bellicose githyanki terrorize the Astral Plane, raiding into other worlds to plunder the multiverse of its magic and riches. The githzerai live apart from the rest of the cosmos, content within the confines of their fortresses floating through the chaos of Limbo. Although the two groups of gith despise each other, their hatred for the mind flayers from whom they escaped endures, and both githyanki and githzerai are dedicated to hunting their ancestral foes.

> [!quote]- A quote from Mordenkainen  
> 
> What would become of this multiverse if githyanki didn't guard the Astral Plane from the illithid menace? What would reality become if beings of thought ruled the plane of thought?


```statblock
"name": "Githzerai Enlightened (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Lawful Neutral"
"ac": !!int "18"
"hp": !!int "112"
"hit_dice": "15d8 + 45"
"stats":
- !!int "14"
- !!int "19"
- !!int "16"
- !!int "17"
- !!int "19"
- !!int "13"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "8"
  "Wisdom": !!int "8"
  "Intelligence": !!int "7"
  "Strength": !!int "6"
"skillsaves":
  "Insight": !!int "8"
  "Perception": !!int "8"
  "Arcana": !!int "7"
"senses": "passive Perception 18"
"languages": "Gith"
"cr": "10"
"traits":
- "desc": "The githzerai's innate spellcasting ability is Wisdom (spell save DC 16,\
    \ dice: d20+8 (+8) to hit with spell attacks). It can innately cast the following\
    \ spells, requiring no components:\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\
    \ (the hand is invisible)\n\n1/day each: [haste](/3-Mechanics/CLI/spells/haste.md),\
    \ [plane shift](/3-Mechanics/CLI/spells/plane-shift.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\
    \n3/day each: [blur](/3-Mechanics/CLI/spells/blur.md), [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md),\
    \ [feather fall](/3-Mechanics/CLI/spells/feather-fall.md), [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [see invisibility](/3-Mechanics/CLI/spells/see-invisibility.md), [shield](/3-Mechanics/CLI/spells/shield.md)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "While the githzerai is wearing no armor and wielding no shield, its AC\
    \ includes its Wisdom modifier."
  "name": "Psychic Defense"
"actions":
- "desc": "The githzerai makes three unarmed strikes."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) bludgeoning damage plus dice:3d8|text(13)\
    \ (3d8) psychic damage."
  "name": "Unarmed Strike"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d8 + 4|text(13) (2d8 + 4) bludgeoning damage plus dice:8d12|text(52)\
    \ (8d12) psychic damage. The target must succeed on a DC 16 Wisdom saving throw\
    \ or move 1 round forward in time. A target moved forward in time vanishes for\
    \ the duration. When the effect ends, the target reappears in the space it left\
    \ or in an unoccupied space nearest to that space if it's occupied."
  "name": "Temporal Strike (Recharge 6)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Githzerai%20Enlightened.webp"
```
^statblock

```encounter-table
name: Githzerai Enlightened
creatures:
 - 1: Githzerai Enlightened
```

## Environment

desert, mountain, urban