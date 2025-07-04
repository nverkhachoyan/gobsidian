---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/8
- monster/size/medium
- monster/type/humanoid
aliases: ["Inquisitor of the Mind Fire"]
NoteIcon: monster
BestiaryType: humanoid
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 248
---
# [Inquisitor of the Mind Fire](3-Mechanics\CLI\bestiary\humanoid/inquisitor-of-the-mind-fire-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 248*  

"Evil lurks everywhere. With our minds, we will unearth it, we will plumb its depths, and we will annihilate it." With those words, the psychically gifted priest Ulmed founded the Ulmist Inquisition, an order of psionic inquisitors that seeks to discover the wickedness hiding in people's souls.

In the days before Count Strahd von Zarovich became the first vampire, Strahd thundered across the lands with Ulmed. Their mission was clear: to destroy the infernal powers that had corrupted the world and to ensure that those powers never rose again. Strahd, Ulmed, and their companions hunted Fiends, Undead, Aberrations, and other supernatural threats and were tireless foes of cults like the priests of Osybus. When Strahd fell into darkness, Ulmed was heartbroken at his friend's transformation and changed the inquisition's mission. Instead of focusing on hunting monsters, it would also hunt the seeds of evil that can corrupt a person.

Ulmed and his friends Cosima, Ansel, and Tristian organized the inquisition into three orders, with each one specializing in a type of psionic power. The Order of Cosima harnessed the Mind Fire—their name for the fire of thought that blazes within each person's mind. They used that power to read thoughts, reshape memories, and dominate the recalcitrant. The inquisitors in the Order of Ansel subjected themselves to harsh asceticism in an effort to use psionic energy to empower their own bodies. They succeeded and became the martial arm of the inquisition, represented by a sword. Finally, the Order of Tristian endeavored to use intellect to alter the environment through telekinetic force, and the order's members became the inquisition's scholars, represented by a tome.

Today the inquisition rules the city of Malitain, a vast city-state to the north of Barovia's original site, and the inquisition sends its members throughout the multiverse, seeking to thwart the work of malevolent cults, otherworldly horrors, and the malice of mortals. The zeal of the inquisitors in this work has caused them to be a source of terror in many communities, where folk fear that an overzealous inquisitor might be as great a monster as the fiends the inquisitors originally hunted.

```statblock
"name": "Inquisitor of the Mind Fire (VRGR)"
"size": "Medium"
"type": "humanoid"
"alignment": "Unaligned"
"ac": !!int "16"
"ac_class": "[breastplate](/3-Mechanics/CLI/items/breastplate.md)"
"hp": !!int "77"
"hit_dice": "14d8 + 14"
"stats":
- !!int "10"
- !!int "14"
- !!int "12"
- !!int "17"
- !!int "16"
- !!int "19"
"speed": "30 ft."
"saves":
  "Charisma": !!int "7"
  "Wisdom": !!int "6"
  "Intelligence": !!int "6"
"skillsaves":
  "Insight": !!int "6"
  "Perception": !!int "6"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "truesight 30 ft., passive Perception 16"
"languages": "any three languages, telepathy 120 ft."
"cr": "8"
"traits":
- "desc": "The inquisitor casts one of the following spells, requiring no components\
    \ and using Charisma as the spellcasting ability (spell save DC 15):\n\nAt will:\
    \ [arcane eye](/3-Mechanics/CLI/spells/arcane-eye.md), [calm emotions](/3-Mechanics/CLI/spells/calm-emotions.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [sending](/3-Mechanics/CLI/spells/sending.md),\
    \ [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\n1/day each: [mass\
    \ suggestion](/3-Mechanics/CLI/spells/mass-suggestion.md), [modify memory](/3-Mechanics/CLI/spells/modify-memory.md)"
  "name": "Innate Spellcasting (Psionics)"
"actions":
- "desc": "The inquisitor attacks twice with its Silver Longsword or uses Mind Fire\
    \ twice."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) if used with two hands, plus dice:4d8|text(18) (4d8) force\
    \ damage."
  "name": "Silver Longsword"
- "desc": "The inquisitor targets one creature it can see within 120 feet of it. The\
    \ target must succeed on a DC 15 Intelligence saving throw or take dice:3d8 +\
    \ 4|text(17) (3d8 + 4) psychic damage and be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)\
    \ until the start of the inquisitor's next turn."
  "name": "Mind Fire"
- "desc": "Each creature of the inquisitor's choice that it can see within 60 feet\
    \ of it must succeed on a DC 15 Wisdom saving throw or be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ until the start of the inquisitor's next turn. On the [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target's turn, the inquisitor can telepathically control the target's move,\
    \ action, or both. When controlled in this way, the target can take only the Attack\
    \ (inquisitor chooses the target) or Dash action."
  "name": "Inquisitor's Command (Recharge 5-6)"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Inquisitor%20of%20the%20Mind%20Fire.webp"
```
^statblock

```encounter-table
name: Inquisitor of the Mind Fire
creatures:
 - 1: Inquisitor of the Mind Fire
```