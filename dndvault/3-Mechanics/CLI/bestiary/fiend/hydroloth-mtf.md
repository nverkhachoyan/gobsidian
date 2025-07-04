---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/9
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Hydroloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 249
---
# [Hydroloth](3-Mechanics\CLI\bestiary\fiend/hydroloth-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 249*  

## Hydroloth

Like the thought-stealing waters of the River Styx they inhabit, hydroloths filch the memories of creatures they attack, stealing away their thoughts for delivery to whatever master they happen to serve. Hydroloths are skilled at finding lost things, especially those that have been swallowed up in the deeps.

For amphibious assaults or underwater conflicts, hydroloths have no equal among yugoloths. They sometimes hire themselves out to attack and scuttle ships and raid coastal settlements.

## Yugoloths

Mercenaries that ply their trade throughout the Lower Planes and in other realms, yugoloths have a reputation for effectiveness that is matched only by their desire for ever more wealth. Although yugoloths aren't especially loyal and typically try to exploit every potential loophole in a contract, they undertake any task for which they are hired, no matter how despicable. Yugoloths come in a wide variety of forms, including those described in the Monster Manual and the six creatures presented here.

```statblock
"name": "Hydroloth (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Neutral Evil"
"ac": !!int "15"
"hp": !!int "135"
"hit_dice": "18d8 + 54"
"stats":
- !!int "12"
- !!int "21"
- !!int "16"
- !!int "19"
- !!int "10"
- !!int "14"
"speed": "20 ft., swim 40 ft."
"skillsaves":
  "Insight": !!int "4"
  "Perception": !!int "4"
"damage_vulnerabilities": "fire"
"damage_resistances": "cold; lightning; bludgeoning, piercing, slashing from nonmagical\
  \ attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 14"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "9"
"traits":
- "desc": "The hydroloth's innate spellcasting ability is Intelligence (spell save\
    \ DC 16). It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only), [water walk](/3-Mechanics/CLI/spells/water-walk.md)\n\n3/day\
    \ each: [control water](/3-Mechanics/CLI/spells/control-water.md), [crown of\
    \ madness](/3-Mechanics/CLI/spells/crown-of-madness.md), [fear](/3-Mechanics/CLI/spells/fear.md),\
    \ [phantasmal killer](/3-Mechanics/CLI/spells/phantasmal-killer.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)"
  "name": "Innate Spellcasting"
- "desc": "The hydroloth can breathe air and water."
  "name": "Amphibious"
- "desc": "The hydroloth has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The hydroloth's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "The hydroloth is immune to the waters of the River Styx as well as any\
    \ effect that would steal or modify its memories or detect or read its thoughts."
  "name": "Secure Memory"
- "desc": "While submerged in liquid, the hydroloth has advantage on attack rolls."
  "name": "Watery Advantage"
"actions":
- "desc": "The hydroloth makes two melee attacks. In place of one of these attacks,\
    \ it can cast one spell that takes 1 action to cast."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 5|text(14) (2d8 + 5) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+9 (+9) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 5|text(16) (2d10 + 5) piercing damage."
  "name": "Bite"
- "desc": "The hydroloth targets one creature it can see within 60 feet of it. The\
    \ target takes dice: 4d6|avg|noform (4d6) psychic damage, and it must make\
    \ a DC 16 Intelligence saving throw. On a successful save, the target becomes\
    \ immune to this hydroloth's Steal Memory for 24 hours. On a failed save, the\
    \ target loses all proficiencies, it can't cast spells, it can't understand language,\
    \ and if its Intelligence and Charisma scores are higher than 5, they become 5.\
    \ Each time the target finishes a long rest, it can repeat the saving throw, ending\
    \ the effect on itself on a success. A [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ or [remove curse](/3-Mechanics/CLI/spells/remove-curse.md) spell cast on the\
    \ target ends this effect early."
  "name": "Steal Memory (1/Day)"
- "desc": "The hydroloth magically teleports, along with any equipment it is wearing\
    \ or carrying, up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Hydroloth.webp"
```
^statblock

```encounter-table
name: Hydroloth
creatures:
 - 1: Hydroloth
```