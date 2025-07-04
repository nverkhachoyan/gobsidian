---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/24
- monster/size/large
- monster/type/fiend/demon
- monster/type/fiend/shapechanger
aliases: ["Graz'zt"]
NoteIcon: monster
BestiaryType: fiend (demon, shapechanger)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 149
---
# [Graz'zt](3-Mechanics\CLI\bestiary\npc/grazzt-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 149*  

## Graz'zt

The appearance of the Dark Prince is a warning that not all beautiful things are good. Standing nearly nine feet tall, Graz'zt strikes the perfect figure of untamed desire, every plane and curve of his body, every glance of his burning eyes, promising a mixture of pleasure and pain. A subtle wrongness pervades his beauty, from the cruel cast of his features to the six fingers on each hand and six toes on each foot. Graz'zt can also transform himself at will, appearing in any humanoid form that pleases him, or his onlookers, all equally tempting in their own ways.

Graz'zt surrounds himself with the finest of things and the most attractive of servants, and he adorns himself in silks and leathers both striking and disturbing in their workmanship. His lair, and those of his cultists, are pleasure palaces where nothing is forbidden, save moderation or kindness.

The dark Prince of Pleasure considers restriction the only sin, and takes what he wants. Cults devoted to him are secret societies of indulgence, often using their debauchery to subjugate others through blackmail, addiction, and manipulation. They frequently wear alabaster masks with ecstatic expressions and ostentatious dress and body ornamentation to their secret assignations.

Although he prefers charm and subtle manipulation, Graz'zt is capable of terrible violence when provoked. He wields the greatsword Angdrelve, also called Wave of Sorrow, its wavy, razor-edged blade dripping acid at his command.

## Graz'zt's Lair

Graz'zt's principal lair is his Argent Palace, a grandiose structure in the city of Zelatar, found within his Abyssal domain of Azzatar. Graz'zt's maddening influence radiates outward in a tangible ripple, warping reality around him. Given enough time in a single location, Graz'zt can twist it with his madness. Graz'zt's lair is a den of ostentation and hedonism. It is adorned with finery and decorations so decadent that even the wealthiest of mortals would blush at the excess. Within Graz'zt's lairs, followers, thralls, and subjects alike are forced to slake Graz'zt's thirst for pageantry and pleasure.

```statblock
"name": "Graz'zt (MTF)"
"size": "Large"
"type": "fiend"
"subtype": "demon, shapechanger"
"alignment": "Chaotic Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "346"
"hit_dice": "33d10 + 165"
"stats":
- !!int "22"
- !!int "15"
- !!int "21"
- !!int "23"
- !!int "21"
- !!int "26"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "9"
  "Wisdom": !!int "12"
  "Constitution": !!int "12"
"skillsaves":
  "Deception": !!int "15"
  "Insight": !!int "12"
  "Perception": !!int "12"
  "Persuasion": !!int "15"
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison; bludgeoning, piercing, slashing that is nonmagical"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 22"
"languages": "all, telepathy 120 ft."
"cr": "24"
"traits":
- "desc": "Graz'zt's spellcasting ability is Charisma (spell save DC 23). He can innately\
    \ cast the following spells, requiring no material components:\n\nAt will:\
    \ [charm person](/3-Mechanics/CLI/spells/charm-person.md), [crown of madness](/3-Mechanics/CLI/spells/crown-of-madness.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [dissonant whispers](/3-Mechanics/CLI/spells/dissonant-whispers.md)\n\n1/day\
    \ each: [dominate monster](/3-Mechanics/CLI/spells/dominate-monster.md), [greater\
    \ invisibility](/3-Mechanics/CLI/spells/greater-invisibility.md)\n\n3/day each:\
    \ [counterspell](/3-Mechanics/CLI/spells/counterspell.md), [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [dominate person](/3-Mechanics/CLI/spells/dominate-person.md), [sanctuary](/3-Mechanics/CLI/spells/sanctuary.md),\
    \ [telekinesis](/3-Mechanics/CLI/spells/telekinesis.md), [teleport](/3-Mechanics/CLI/spells/teleport.md)"
  "name": "Innate Spellcasting"
- "desc": "Graz'zt can use his action to polymorph into a form that resembles a Medium\
    \ humanoid, or back into his true form. Aside from his size, his statistics are\
    \ the same in each form. Any equipment he is wearing or carrying isn't transformed."
  "name": "Shapechanger"
- "desc": "If Graz'zt fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Graz'zt has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Graz'zt's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "Graz'zt attacks twice with Wave of Sorrow."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 10 ft., one\
    \ target. Hit: dice:4d6 + 6|text(20) (4d6 + 6) slashing damage plus dice:3d6|text(10)\
    \ (3d6) acid damage."
  "name": "Wave of Sorrow (Greatsword)"
- "desc": "Graz'zt magically teleports, along with any equipment he is wearing or\
    \ carrying, up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Graz'zt attacks once with Wave of Sorrow."
  "name": "Attack"
- "desc": "One creature [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by\
    \ Graz'zt that Graz'zt can see must use its reaction to move up to its speed as\
    \ Graz'zt directs."
  "name": "Dance, My Puppet"
- "desc": "Graz'zt casts [crown of madness](/3-Mechanics/CLI/spells/crown-of-madness.md)\
    \ or dissonant whispers."
  "name": "Sow Discord"
- "desc": "Graz'zt uses his Teleport action."
  "name": "Teleport"
"lair_actions":
- "desc": "On Initiative count 20 (losing initiative ties), Graz'zt can take a lair\
    \ action to cause one of the following effects; he can't use the same effect two\
    \ rounds in a row:"
  "name": ""
- "desc": "- Graz'zt casts the [command](/3-Mechanics/CLI/spells/command.md) spell\
    \ on every creature of his choice in the lair. He needn't see each one, but he\
    \ must be aware that an individual is in the lair to target that creature. He\
    \ issues the same command to all the targets.  \n- Smooth surfaces within the\
    \ lair become as reflective as a polished mirror. Until a different lair action\
    \ is used, creatures within the lair have disadvantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Graz'zt's lair is warped by his magic, creating one\
    \ or more of the following effects:"
  "name": ""
- "desc": "- Flat surfaces within 1 mile of the lair that are made of stone or metal\
    \ become highly reflective, as though polished to a shine. These surfaces become\
    \ supernaturally mirrorlike.  \n- Wild beasts within 6 miles of the lair break\
    \ into frequent conflicts and coupling, mirroring the behavior that occurs during\
    \ their mating seasons.  \n- If a humanoid spends at least 1 hour within 1 mile\
    \ of the lair, that creature must succeed on a DC 23 Wisdom saving throw or descend\
    \ into a madness determined by the Madness of Graz'zt table. A creature that succeeds\
    \ on this saving throw can't be affected by this regional effect again for 24\
    \ hours.  "
  "name": ""
- "desc": "If Graz'zt dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
- "desc": "If a creature goes mad in Graz'zt's lair or within line of sight of the\
    \ demon lord, roll on the Madness of Graz'zt table to determine the nature of\
    \ the madness, which is a character flaw that lasts until cured. See the \"Dungeon\
    \ Master's Guide\" for more on madness.\n\nMadness of Graz'zt\n\ndice: [](grazzt-mtf.md#^madness-of-grazzt)\n\
    \n| dice: d100 | Flaw (lasts until cured) |\n|------------|--------------------------|\n\
    | 01–20 | \"There is nothing in the world more important than me and my desires.\"\
    \ |\n| 21–40 | \"Anyone who doesn't do exactly what I say doesn't deserve to live.\"\
    \ |\n| 41–60 | \"Mine is the path of redemption. Anyone who says otherwise is\
    \ intentionally misleading you.\" |\n| 61–80 | \"I will not rest until I have\
    \ made someone else mine, and doing so is more important to me than my own life—\
    or the lives of others.\" |\n| 81–90 | \"My own pleasure is of paramount importance.\
    \ Everything else, including social graces, is a triviality.\" |\n| 91–00 | \"\
    Anything that can bring me happiness should be enjoyed immediately. There is no\
    \ point to saving anything pleasurable for later.\" |\n^madness-of-grazzt"
  "name": "Madness of Graz'zt"
"source":
- "MTF"
- "OotA"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Graz%27zt.webp"
```
^statblock

```encounter-table
name: Graz'zt
creatures:
 - 1: Graz'zt
```