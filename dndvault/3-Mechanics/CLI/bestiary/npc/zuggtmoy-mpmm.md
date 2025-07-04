---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/23
- monster/size/large
- monster/type/fiend/demon
aliases: ["Zuggtmoy"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 281, Mordenkainen's Tome of Foes p. 157
---
# [Zuggtmoy](3-Mechanics\CLI\bestiary\npc/zuggtmoy-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 281, Mordenkainen's Tome of Foes p. 157*  

```statblock
"name": "Zuggtmoy (MPMM)"
"size": "Large"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "304"
"hit_dice": "32d10 + 128"
"stats":
- !!int "22"
- !!int "15"
- !!int "18"
- !!int "20"
- !!int "19"
- !!int "24"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "11"
  "Constitution": !!int "11"
"skillsaves":
  "Perception": !!int "11"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 21"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Zuggtmoy casts one of the following spells, requiring no material components\
    \ and using Charisma as the spellcasting ability (spell save DC 22):\n\nAt will:\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [locate animals or\
    \ plants](/3-Mechanics/CLI/spells/locate-animals-or-plants.md)\n\n1/day each:\
    \ [etherealness](/3-Mechanics/CLI/spells/etherealness.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)\n\
    \n3/day each: [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [entangle](/3-Mechanics/CLI/spells/entangle.md),\
    \ [plant growth](/3-Mechanics/CLI/spells/plant-growth.md)"
  "name": "Spellcasting"
- "desc": "If Zuggtmoy fails a saving throw, she can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Zuggtmoy has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "Zuggtmoy makes three Pseudopod attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d8 + 6|text(15) (2d8 + 6) force damage plus dice:2d8|text(9)\
    \ (2d8) poison damage."
  "name": "Pseudopod"
"bonus_actions":
- "desc": "Zuggtmoy releases spores that burst out in a cloud that fills a 20-foot-radius\
    \ sphere centered on her, and it lingers for 1 minute. Any creature in the cloud\
    \ when it appears, or that enters it later, must make a DC 19 Constitution saving\
    \ throw. On a successful save, the creature can't be infected by these spores\
    \ for 24 hours. On a failed save, the creature is infected with a disease called\
    \ the spores of Zuggtmoy, which lasts until the creature is cured of the disease\
    \ or dies. While infected in this way, the creature can't be reinfected, and it\
    \ must repeat the saving throw at the end of every 24 hours, ending the infection\
    \ on a success. On a failure, the infected creature's body is slowly taken over\
    \ by fungal growth, and after three such failed saves, the creature dies and is\
    \ reanimated as a [spore servant](/3-Mechanics/CLI/bestiary/plant/quaggoth-spore-servant.md)\
    \ if it's a type of creature that can be."
  "name": "Infestation Spores (3/Day)"
- "desc": "Zuggtmoy releases spores that burst out in a cloud that fills a 20-foot-radius\
    \ sphere centered on her, and it lingers for 1 minute. Humanoids and Beasts in\
    \ the cloud when it appears, or that enter it later, must make a DC 19 Wisdom\
    \ saving throw. On a successful save, a creature can't be infected by these spores\
    \ for 24 hours. On a failed save, the creature is infected with a disease called\
    \ the influence of Zuggtmoy for 24 hours. While infected in this way, the creature\
    \ is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by her and can't\
    \ be reinfected by these spores."
  "name": "Mind Control Spores (Recharge 5-6)"
"reactions":
- "desc": "When Zuggtmoy is hit by an attack roll, one creature within 10 feet of\
    \ her that is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by her is\
    \ hit by the attack instead."
  "name": "Protective Thrall"
"legendary_actions":
- "desc": "Zuggtmoy makes one Pseudopod attack."
  "name": "Attack"
- "desc": "One creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by\
    \ Zuggtmoy that she can see must use its reaction, if a available, to move up\
    \ to its speed as she directs or to make one weapon attack against a target that\
    \ she designates."
  "name": "Exert Will"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Zuggtmoy.webp"
```
^statblock

```encounter-table
name: Zuggtmoy
creatures:
 - 1: Zuggtmoy
```