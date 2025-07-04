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
aliases: ["Githyanki Gish"]
NoteIcon: monster
BestiaryType: humanoid (gith)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 205
---
# [Githyanki Gish](3-Mechanics\CLI\bestiary\humanoid/githyanki-gish-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 205*  

> [!quote]- A quote from Mordenkainen  
> 
> What would become of this multiverse if githyanki didn't guard the Astral Plane from the illithid menace? What would reality become if beings of thought ruled the plane of thought?

## Githyanki Gish

Their keen minds and psionic gifts allow the githyanki to master magic. Gish blend their magical abilities with swordplay to become dangerous foes in battle. Their specialized capabilities make them well suited for assassination, raiding, and espionage.

## Gith

The descendants of an ancient people—so old their original name has been lost—have turned against each other, becoming vicious enemies divided over mortality, purpose, and the machinations of their leaders. The bellicose githyanki terrorize the Astral Plane, raiding into other worlds to plunder the multiverse of its magic and riches. The githzerai live apart from the rest of the cosmos, content within the confines of their fortresses floating through the chaos of Limbo. Although the two groups of gith despise each other, their hatred for the mind flayers from whom they escaped endures, and both githyanki and githzerai are dedicated to hunting their ancestral foes.

```statblock
"name": "Githyanki Gish (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "gith"
"alignment": "Lawful Evil"
"ac": !!int "17"
"ac_class": "[half plate armor](/3-Mechanics/CLI/items/half-plate-armor.md)"
"hp": !!int "123"
"hit_dice": "19d8 + 38"
"stats":
- !!int "17"
- !!int "15"
- !!int "14"
- !!int "16"
- !!int "15"
- !!int "16"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "6"
  "Intelligence": !!int "7"
  "Constitution": !!int "6"
"skillsaves":
  "Stealth": !!int "6"
  "Insight": !!int "6"
  "Perception": !!int "6"
"senses": "passive Perception 16"
"languages": "Gith"
"cr": "10"
"traits":
- "desc": "The githyanki's innate spellcasting ability is Intelligence (spell save\
    \ DC 15, dice: d20+7 (+7) to hit with spell attacks). It can innately cast\
    \ the following spells, requiring no components:\n\nAt will: [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\
    \ (the hand is invisible)\n\n1/day each: [plane shift](/3-Mechanics/CLI/spells/plane-shift.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)\n\n3/day each: [jump](/3-Mechanics/CLI/spells/jump.md),\
    \ [misty step](/3-Mechanics/CLI/spells/misty-step.md), [nondetection](/3-Mechanics/CLI/spells/nondetection.md)\
    \ (self only)"
  "name": "Innate Spellcasting (Psionics)"
- "desc": "The githyanki is an 8th-level spellcaster. Its spellcasting ability is\
    \ Intelligence (spell save DC 15, dice: d20+7 (+7) to hit with spell attacks).\
    \ The githyanki has the following wizard spells prepared:\n\nCantrips (at will):\
    \ [blade ward](/3-Mechanics/CLI/spells/blade-ward.md), [light](/3-Mechanics/CLI/spells/light.md),\
    \ [message](/3-Mechanics/CLI/spells/message.md), [true strike](/3-Mechanics/CLI/spells/true-strike.md)\n\
    \n1st level (4 slots): [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md),\
    \ [magic missile](/3-Mechanics/CLI/spells/magic-missile.md), [sleep](/3-Mechanics/CLI/spells/sleep.md),\
    \ [thunderwave](/3-Mechanics/CLI/spells/thunderwave.md)\n\n2nd level (3 slots):\
    \ [blur](/3-Mechanics/CLI/spells/blur.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md),\
    \ [levitate](/3-Mechanics/CLI/spells/levitate.md)\n\n3rd level (3 slots):\
    \ [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [fireball](/3-Mechanics/CLI/spells/fireball.md),\
    \ [haste](/3-Mechanics/CLI/spells/haste.md)\n\n4th level (2 slots): [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md)"
  "name": "Spellcasting"
- "desc": "When the githyanki uses its action to cast a spell, it can make one weapon\
    \ attack as a bonus action."
  "name": "War Magic"
"actions":
- "desc": "The githyanki makes two longsword attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) slashing damage plus dice:4d8|text(18)\
    \ (4d8) psychic damage, or dice:1d10 + 3|text(8) (1d10 + 3) slashing damage\
    \ plus dice:4d8|text(18) (4d8) psychic damage if used with two hands."
  "name": "Longsword"
"source":
- "MTF"
- "WDMM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Githyanki%20Gish.webp"
```
^statblock

```encounter-table
name: Githyanki Gish
creatures:
 - 1: Githyanki Gish
```

## Environment

desert, mountain, urban