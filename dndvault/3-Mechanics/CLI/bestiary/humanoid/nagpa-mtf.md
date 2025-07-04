---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/17
- monster/environment/coastal
- monster/environment/desert
- monster/environment/forest
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/nagpa
aliases: ["Nagpa"]
NoteIcon: monster
BestiaryType: humanoid (nagpa)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 215
---
# [Nagpa](3-Mechanics\CLI\bestiary\humanoid/nagpa-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 215*  

> [!quote]- A quote from Mordenkainen  
> 
> Nagpas can't learn from that which isn't destroyed. Imagine how bewildering and terrifying a city full of life must be to them. Is it any wonder that they seek the ruin of such places?

## Nagpa

Long ago, the Raven Queen cursed a cabal of thirteen powerful wizards for meddling in a ritual that would have helped avert a war between the gods. She stripped them of their beauty, turning them into scabrous, birdlike monstrosities. The nagpas now plot as they ever did, but they now strive to bring about terrible, world-shaking calamities so they can pry secrets and power from the wreckage their conspiracies create.

The nagpas fear the Raven Queen and do their best to avoid her and her agents. When it's impossible to do so, they become cringing, fawning things, eager to please and thereby escape the cold gaze of the being who brought them so low. All of the original thirteen remain alive, thanks to their cunning and their willingness to do whatever is necessary to survive.

## Looters of Civilization

The curse the Raven Queen placed on the nagpas restricts the ways in which they can acquire new lore and magical power, barring them from any source except for the ruins left behind from fallen civilizations and great calamities. For this reason, nagpas turn their efforts to bringing about such ends, so they can loot the libraries, plunder the vaults, and gather up secrets of arcane lore from the wreckage.

## Puppet Masters

Nagpas work in the shadows, manipulating events to bring about ruin. As accomplished magic-users, they can bring to bear an array of spells to make agents of other creatures, influencing their decisions in subtle ways and pulling on strings to make them into unwitting accomplices in their own destruction. Nagpas show great patience in their plots and have several schemes working simultaneously, each at different stages of completion, so if one plan goes awry, they can shift their focus to another. Typically, nagpas only show their handiwork and emerge from the shadows when they can deliver a finishing blow and then revel in the grand devastation their plotting brought about.

```statblock
"name": "Nagpa (MTF)"
"size": "Medium"
"type": "humanoid"
"subtype": "nagpa"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "187"
"hit_dice": "34d8 + 34"
"stats":
- !!int "9"
- !!int "15"
- !!int "12"
- !!int "23"
- !!int "18"
- !!int "21"
"speed": "30 ft."
"saves":
  "Charisma": !!int "11"
  "Wisdom": !!int "10"
  "Intelligence": !!int "12"
"skillsaves":
  "Deception": !!int "11"
  "Insight": !!int "10"
  "Perception": !!int "10"
  "History": !!int "12"
  "Arcana": !!int "12"
"senses": "truesight 120 ft., passive Perception 20"
"languages": "Common plus up to five other languages"
"cr": "17"
"traits":
- "desc": "The nagpa is a 15th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (spell save DC 20, dice: d20+12 (+12) to hit with spell attacks). A nagpa\
    \ has the following wizard spells prepared:\n\nCantrips (at will): [chill\
    \ touch](/3-Mechanics/CLI/spells/chill-touch.md), [fire bolt](/3-Mechanics/CLI/spells/fire-bolt.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [message](/3-Mechanics/CLI/spells/message.md),\
    \ [minor illusion](/3-Mechanics/CLI/spells/minor-illusion.md)\n\n1st level (4\
    \ slots): [charm person](/3-Mechanics/CLI/spells/charm-person.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [protection from evil and good](/3-Mechanics/CLI/spells/protection-from-evil-and-good.md),\
    \ [witch bolt](/3-Mechanics/CLI/spells/witch-bolt.md)\n\n2nd level (3 slots):\
    \ [hold person](/3-Mechanics/CLI/spells/hold-person.md), [ray of enfeeblement](/3-Mechanics/CLI/spells/ray-of-enfeeblement.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n3rd level (3 slots):\
    \ [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [fireball](/3-Mechanics/CLI/spells/fireball.md),\
    \ [fly](/3-Mechanics/CLI/spells/fly.md)\n\n4th level (3 slots): [confusion](/3-Mechanics/CLI/spells/confusion.md),\
    \ [hallucinatory terrain](/3-Mechanics/CLI/spells/hallucinatory-terrain.md), [wall\
    \ of fire](/3-Mechanics/CLI/spells/wall-of-fire.md)\n\n5th level (2 slots):\
    \ [dominate person](/3-Mechanics/CLI/spells/dominate-person.md), [dream](/3-Mechanics/CLI/spells/dream.md),\
    \ [geas](/3-Mechanics/CLI/spells/geas.md)\n\n6th level (1 slots): [circle\
    \ of death](/3-Mechanics/CLI/spells/circle-of-death.md), [disintegrate](/3-Mechanics/CLI/spells/disintegrate.md)\n\
    \n7th level (1 slots): [etherealness](/3-Mechanics/CLI/spells/etherealness.md),\
    \ [prismatic spray](/3-Mechanics/CLI/spells/prismatic-spray.md)\n\n8th level\
    \ (1 slots): [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md)"
  "name": "Spellcasting"
- "desc": "As a bonus action, the nagpa targets one creature it can see within 90\
    \ feet of it. The target must make a DC 20 Charisma saving throw. An evil creature\
    \ makes the save with disadvantage. On a failed save, the target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the nagpa until the start of the nagpa's next turn. On a successful save,\
    \ the target becomes immune to the nagpa's Corruption for the next 24 hours."
  "name": "Corruption"
- "desc": "As a bonus action, the nagpa forces each creature within 30 feet of it\
    \ to succeed on a DC 20 Wisdom saving throw or be [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ for 1 minute. A [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed)\
    \ target can repeat the saving throw at the end of each of its turns, ending the\
    \ effect on itself on a success. Undead and constructs are immune to this effect."
  "name": "Paralysis (Recharge 6)"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) bludgeoning damage."
  "name": "Staff"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Nagpa.webp"
```
^statblock

```encounter-table
name: Nagpa
creatures:
 - 1: Nagpa
```

## Environment

coastal, desert, forest, swamp, underdark, urban