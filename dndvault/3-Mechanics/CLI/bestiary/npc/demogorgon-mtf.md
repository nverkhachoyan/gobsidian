---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/26
- monster/size/huge
- monster/type/fiend/demon
aliases: ["Demogorgon"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 144
---
# [Demogorgon](3-Mechanics\CLI\bestiary\npc/demogorgon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 144*  

## Demogorgon

Prince of Demons, the Sibilant Beast, and Master of the Spiraling Depths, Demogorgon is the embodiment of chaos, madness, and destruction, seeking to corrupt all that is good and undermine order in the multiverse, to see everything dragged howling into the infinite depths of the Abyss.

The demon lord is a meld of different forms, with a saurian lower body and clawed, webbed feet, as well as suckered tentacles sprouting from the shoulders of a great apelike torso, surmounted by two hideous simian heads, named Aameul and Hathradiah, both equally mad. Their gaze brings madness and confusion to any who confront it.

Similarly, the spiraling Y sign of Demogorgon's cult can inspire madness in those who contemplate it for too long. All the followers of the Prince of Demons go mad, sooner or later.

## Demogorgon's Lair

Demogorgon makes his lair in a palace called Abysm, found on a layer of the Abyss known as the Gaping Maw. Demogorgon's lair is a place of madness and duality; the portion of the palace that lies above water takes the form of two serpentine towers, each crowned by a skull-shaped minaret. There, Demogorgon's heads contemplate the mysteries of the arcane while arguing about how best to obliterate their rivals. The bulk of this palace extends deep underwater, in chill and darkened caverns.

```statblock
"name": "Demogorgon (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "22"
"ac_class": "natural armor"
"hp": !!int "406"
"hit_dice": "28d12 + 224"
"stats":
- !!int "29"
- !!int "14"
- !!int "26"
- !!int "20"
- !!int "17"
- !!int "25"
"speed": "50 ft., swim 50 ft."
"saves":
  "Charisma": !!int "15"
  "Dexterity": !!int "10"
  "Wisdom": !!int "11"
  "Constitution": !!int "16"
"skillsaves":
  "Insight": !!int "11"
  "Perception": !!int "19"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 29"
"languages": "all, telepathy 120 ft."
"cr": "26"
"traits":
- "desc": "Demogorgon's spellcasting ability is Charisma (spell save DC 23). Demogorgon\
    \ can innately cast the following spells, requiring no material components:\n\n\
    At will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [major image](/3-Mechanics/CLI/spells/major-image.md)\n\
    \n1/day each: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md), [project\
    \ image](/3-Mechanics/CLI/spells/project-image.md)\n\n3/day each: [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md)"
  "name": "Innate Spellcasting"
- "desc": "If Demogorgon fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Demogorgon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Demogorgon's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Demogorgon has advantage on saving throws against being [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned),\
    \ or knocked [unconscious](/3-Mechanics/CLI/rules/conditions.md#unconscious)."
  "name": "Two Heads"
"actions":
- "desc": "Demogorgon makes two tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 10 ft., one\
    \ target. Hit: dice:3d12 + 9|text(28) (3d12 + 9) bludgeoning damage. If\
    \ the target is a creature, it must succeed on a DC 23 Constitution saving throw\
    \ or its hit point maximum is reduced by an amount equal to the damage taken.\
    \ This reduction lasts until the target finishes a long rest. The target dies\
    \ if its hit point maximum is reduced to 0."
  "name": "Tentacle"
- "desc": "Demogorgon turns his magical gaze toward one creature that he can see within\
    \ 120 feet of him. That target must make a DC 23 Wisdom saving throw. Unless the\
    \ target is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ it can avert its eyes to avoid the gaze and to automatically succeed on the\
    \ save. If the target does so, it can't see Demogorgon until the start of his\
    \ next turn. If the target looks at him in the meantime, it must immediately make\
    \ the save.\n\nIf the target fails the save, the target suffers one of the following\
    \ effects of Demogorgon's choice or at random:\n\n1. Beguiling Gaze. The target\
    \ is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until the start of\
    \ Demogorgon's next turn or until Demogorgon is no longer within line of sight.\n\
    \n2. Hypnotic Gaze. The target is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by Demogorgon until the start of Demogorgon's next turn. Demogorgon chooses\
    \ how the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) target uses\
    \ its actions, reactions, and movement. Because this gaze requires Demogorgon\
    \ to focus both heads on the target, he can't use his Maddening Gaze legendary\
    \ action until the start of his next turn.\n\n3. Insanity Gaze. The target suffers\
    \ the effect of the [confusion](/3-Mechanics/CLI/spells/confusion.md) spell without\
    \ making a saving throw. The effect lasts until the start of Demogorgon's next\
    \ turn. Demogorgon doesn't need to concentrate on the spell."
  "name": "Gaze"
"legendary_actions":
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d10 + 9|text(20) (2d10 + 9) bludgeoning damage plus\
    \ dice:2d10|text(11) (2d10) necrotic damage."
  "name": "Tail"
- "desc": "Demogorgon uses his Gaze action, and must choose either the Beguiling Gaze\
    \ or the Insanity Gaze effect."
  "name": "Maddening Gaze"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Demogorgon can take a\
    \ lair action to cause one of the following effects: Demogorgon can't use the\
    \ same effect two rounds in a row:"
  "name": ""
- "desc": "- Demogorgon creates an illusory duplicate of himself, which appears in\
    \ his own space and lasts until initiative count 20 of the next round. On his\
    \ turn, Demogorgon can move the illusory duplicate a distance equal to his walking\
    \ speed (no action required). The first time a creature or object interacts physically\
    \ with Demogorgon (for example, hitting him with an attack), there is a 50% chance\
    \ chance that it is the illusory duplicate that is being affected, not Demogorgon\
    \ himself, in which case the illusion disappears.  \n- Demogorgon casts the [darkness](/3-Mechanics/CLI/spells/darkness.md)\
    \ spell four times at its lowest level, targeting different areas with the spell.\
    \ Demogorgon doesn't need to concentrate on the spells, which end on initiative\
    \ count 20 of the next round.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Demogorgon's lair is warped by his magic, creating\
    \ one or more of the following effects:"
  "name": ""
- "desc": "- The area within 6 miles of the lair becomes overpopulated with lizards,\
    \ poisonous snakes, and other venomous beasts.  \n- Beasts within 1 mile of the\
    \ lair become violent and crazed-even creatures that are normally docile.  \n\
    - If a humanoid spends at least 1 hour within 1 mile of the lair, that creature\
    \ must succeed on a DC 23 Wisdom saving throw or descend into a madness determined\
    \ by the Madness of Demogorgon table. A creature that succeeds on this saving\
    \ throw can't be affected by this regional effect again for 24 hours.  "
  "name": ""
- "desc": "If Demogorgon dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Demogorgon's lair or within line of sight of\
    \ the demon lord, roll on the Madness of Demogorgon table to determine the nature\
    \ of the madness, which is a character flaw that lasts until cured. See the \"\
    Dungeon Master's Guide\" for more on madness.\n\nMadness of Demogorgon\n\n\
    dice: [](demogorgon-mtf.md#^madness-of-demogorgon)\n\n| dice: d100 | Flaw (lasts\
    \ until cured) |\n|------------|--------------------------|\n| 01–20 | \"Someone\
    \ is plotting to kill me. I need to strike first to stop them!\" |\n| 21–40 |\
    \ \"There is only one solution to my problems: kill them all!\" |\n| 41–60 | \"\
    There is more than one mind inside my head.\" |\n| 61–80 | \"If you don't agree\
    \ with me, I'll beat you into submission to get my way.\" |\n| 81–00 | \"I can't\
    \ allow anyone to touch anything that belongs to me. They might try to take it\
    \ away from me!\" |\n^madness-of-demogorgon"
  "name": "Madness of Demogorgon"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Demogorgon.webp"
```
^statblock

```encounter-table
name: Demogorgon
creatures:
 - 1: Demogorgon
```