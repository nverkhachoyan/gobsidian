---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/23
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Juiblex"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 160, Mordenkainen's Tome of Foes p. 151
---
# [Juiblex](3-Mechanics\CLI\bestiary\npc/juiblex-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 160, Mordenkainen's Tome of Foes p. 151*  

```statblock
"name": "Juiblex (MPMM)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "18"
"ac_class": "natural armor"
"hp": !!int "350"
"hit_dice": "28d12 + 168"
"stats":
- !!int "24"
- !!int "10"
- !!int "23"
- !!int "20"
- !!int "20"
- !!int "16"
"speed": "30 ft., climb 30 ft."
"saves":
  "Dexterity": !!int "7"
  "Wisdom": !!int "12"
  "Constitution": !!int "13"
"skillsaves":
  "Perception": !!int "12"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "acid; poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
  \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
  \ [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned), [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)"
"senses": "truesight 120 ft., passive Perception 22"
"languages": "all, telepathy 120 ft."
"cr": "23"
"traits":
- "desc": "Juiblex casts one of the following spells, requiring no material components\
    \ and using Wisdom as the spellcasting ability (spell save DC 20):\n\nAt will:\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md)\n\n3/day each: [contagion](/3-Mechanics/CLI/spells/contagion.md),\
    \ [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md)"
  "name": "Spellcasting"
- "desc": "Any creature other than an Ooze that starts its turn within 10 feet of\
    \ Juiblex must succeed on a DC 21 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until the start of the creature's next turn."
  "name": "Foul"
- "desc": "If Juiblex fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Juiblex has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Juiblex regains 20 hit points at the start of its turn. If it takes fire\
    \ or radiant damage, this trait doesn't function at the start of its next turn.\
    \ Juiblex dies only if it starts its turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
- "desc": "Juiblex can climb difficult surfaces, including upside down on ceilings,\
    \ without needing to make an ability check."
  "name": "Spider Climb"
"actions":
- "desc": "Juiblex makes three Acid Lash attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+14 (+14) to hit, reach 10\
    \ ft. or range 60/120 ft., one target. Hit: dice:4d6 + 7|text(21) (4d6 +\
    \ 7) acid damage. Any creature killed by this attack is drawn into Juiblex's\
    \ body, where the corpse is dissolved after 1 minute."
  "name": "Acid Lash"
- "desc": "Juiblex spews out a corrosive slime, targeting one creature that it can\
    \ see within 60 feet of it. The target must succeed on a DC 21 Dexterity saving\
    \ throw or take dice:10d10|text(55) (10d10) acid damage. Unless the target\
    \ avoids taking all of this damage, any metal armor worn by the target takes a\
    \ permanent −1 penalty to the AC it offers, and any metal weapon the target is\
    \ carrying or wearing takes a permanent −1 penalty to damage rolls. The penalty\
    \ worsens each time a target is subjected to this effect. If the penalty on an\
    \ object drops to −5, the object is destroyed. The penalty on an object can be\
    \ removed by the [mending](/3-Mechanics/CLI/spells/mending.md) spell."
  "name": "Eject Slime (Recharge 5-6)"
"legendary_actions":
- "desc": "Juiblex makes one Acid Lash attack."
  "name": "Attack"
- "desc": "Melee Weapon Attack: dice: d20+14 (+14) to hit, reach 10 ft., one\
    \ creature. Hit: dice:4d6 + 7|text(21) (4d6 + 7) poison damage, and the\
    \ target is slimed. Until the slime is scraped off with an action, the target\
    \ is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), and any creature,\
    \ other than an Ooze, is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ while within 10 feet of the target."
  "name": "Corrupting Touch (Costs 2 Actions)"
"source":
- "MPMM"
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Juiblex.webp"
```
^statblock

```encounter-table
name: Juiblex
creatures:
 - 1: Juiblex
```