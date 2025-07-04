---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/23
- monster/size/huge
- monster/type/celestial/titan
aliases: ["Empyrean"]
NoteIcon: monster
BestiaryType: celestial (titan)
SourceType: Bestiary
BookSource: Monster Manual p. 130
---
# [Empyrean](3-Mechanics\CLI\bestiary\celestial/empyrean.md)
*Source: Monster Manual p. 130*  

Empyreans are the celestial children of the gods of the Upper Planes. They are universally beautiful, statuesque, and self-assured.

## Manifest Emotion

An empyrean can experience deity-like fits of serenity or rage. It can affect the environment around it by its mood. When an empyrean is unhappy, the clouds might cry tears of salt water, the wildflowers in surrounding meadows might wilt, dead fish might wash ashore in lakes or rivers, or a nearby forest might lose the leaves from its trees. When an empyrean is jubilant, sunlight follows it everywhere, small animals frolic in its footsteps, and birds fill the sky with their pleasing songs.

## Evil Empyreans

A few empyreans have turned to evil after venturing to the Lower Planes and becoming corrupted, or as the result of being cursed by evil gods. An evil empyrean can't survive long on the Upper Planes and usually retreats to the Material Plane, where it can rule over a kingdom of mortals as an indomitable tyrant.

## Immortal Titans

Empyreans don't age but can be slain. Because few empyreans can imagine their own demise, they fight fearlessly when drawn into battle, refusing to believe that the end is upon them even when standing at death's door. When an empyrean dies, its spirit returns to its home plane. There, one of the fallen empyrean's parents resurrects the empyrean unless he or she has a good reason not to.

```statblock
"name": "Empyrean"
"size": "Huge"
"type": "celestial"
"subtype": "titan"
"alignment": "Chaotic Good or Neutral Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "313"
"hit_dice": "19d12 + 190"
"stats":
- !!int "30"
- !!int "21"
- !!int "30"
- !!int "21"
- !!int "22"
- !!int "27"
"speed": "50 ft., fly 50 ft., swim 50 ft."
"saves":
  "Charisma": !!int "15"
  "Wisdom": !!int "13"
  "Intelligence": !!int "12"
  "Strength": !!int "17"
"skillsaves":
  "Insight": !!int "13"
  "Persuasion": !!int "15"
"damage_immunities": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "truesight 120 ft., passive Perception 16"
"languages": "all"
"cr": "23"
"traits":
- "desc": "The empyrean's innate spellcasting ability is Charisma (spell save DC 23,\
    \ dice: d20+15 (+15) to hit with spell attacks). It can innately cast the\
    \ following spells, requiring no material components:\n\nAt will: [greater\
    \ restoration](/3-Mechanics/CLI/spells/greater-restoration.md), [pass without\
    \ trace](/3-Mechanics/CLI/spells/pass-without-trace.md), [water breathing](/3-Mechanics/CLI/spells/water-breathing.md),\
    \ [water walk](/3-Mechanics/CLI/spells/water-walk.md)\n\n1/day each: [commune](/3-Mechanics/CLI/spells/commune.md),\
    \ [dispel evil and good](/3-Mechanics/CLI/spells/dispel-evil-and-good.md), [earthquake](/3-Mechanics/CLI/spells/earthquake.md),\
    \ [fire storm](/3-Mechanics/CLI/spells/fire-storm.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\
    \ (self only)"
  "name": "Innate Spellcasting"
- "desc": "If the empyrean fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The empyrean has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The empyrean's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:6d6 + 10|text(31) (6d6 + 10) bludgeoning damage. If\
    \ the target is a creature, it must succeed on a DC 15 Constitution saving throw\
    \ or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until the end\
    \ of the empyrean's next turn."
  "name": "Maul"
- "desc": "Ranged Spell Attack: dice: d20+15 (+15) to hit, range 600 ft., one\
    \ target. Hit: dice:7d6|text(24) (7d6) damage of one of the following types\
    \ (empyrean's choice): acid, cold, fire, force, lightning, radiant, or thunder."
  "name": "Bolt"
"legendary_actions":
- "desc": "The empyrean makes one attack."
  "name": "Attack"
- "desc": "The empyrean bolsters all nonhostile creatures within 120 feet of it until\
    \ the end of its next turn. Bolstered creatures can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and they gain\
    \ advantage on ability checks and saving throws until the end of the empyrean's\
    \ next turn."
  "name": "Bolster"
- "desc": "The empyrean strikes the ground with its maul, triggering an earth tremor.\
    \ All other creatures on the ground within 60 feet of the empyrean must succeed\
    \ on a DC 25 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Trembling Strike (Costs 2 Actions)"
"source":
- "MM"
- "TftYP"
- "WDMM"
- "BGDIA"
- "MOT"
- "SatO"
- "ToFW"
- "BMT"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Empyrean.webp"
```
^statblock

```encounter-table
name: Empyrean
creatures:
 - 1: Empyrean
```