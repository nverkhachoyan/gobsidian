---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/6
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Bodak"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 127
---
# [Bodak](3-Mechanics\CLI\bestiary\undead/bodak-vgm.md)
*Source: Volo's Guide to Monsters p. 127*  

A bodak is the undead remains of someone who revered Orcus. Devoid of life and soul, it exists only to cause death.

## Marked by Orcus

A worshiper of Orcus can take ritual vows while carving the demon lord's symbol on its chest over the heart. Orcus's power flays body, mind, and soul, leaving behind a sentient husk that sucks in all life energy near it. Most bodaks come into being in this way, then unleashed to spread death in Orcus's name.

Orcus created the first bodaks in the Abyss from seven devotees, called the Hierophants of Annihilation. These figures, as mighty as balors, have free will but serve the Prince of Undeath directly. Any one of these bodaks can turn a slain mortal into a bodak with its gaze. Like each Hierophant of Annihilation, every bodak bears the mark of Orcus as a chest wound, an opening where a mortal humanoid's heart would be.

Orcus can recall anything a bodak sees or hears. If he so chooses, he can speak through a bodak to address his enemies and followers directly. Bodaks are extensions of Orcus's will outside the Abyss, serving the demon prince's aims and other minions.

Even nature despises bodaks. The sun burns away a bodak's tainted flesh. The creature's gaze lays waste to the living. Anyone a bodak slays with its gaze withers, its face frozen in a mask of terror. The monster's mere presence is so unnatural that it chills the soul. Animals untrained for war instinctively flee just before a bodak arrives.

## Unhallowed Fragments

A bodak retains vague impressions of its past life. It seeks out both its former allies and its former enemies to destroy them, as its warped soul seeks to erase anything connected to its former life. Minions of Orcus are the one exception to this compulsion; a bodak recognizes them as kindred souls and spares them from its wrath. Anyone who knew the individual before its transformation into a bodak can recognize mannerisms or other subtle clues to its original identity.

## Ravaged Soul

The soul of a creature that becomes a bodak is so damaged that it is unfit for most forms of magical resurrection. Only a [wish](/3-Mechanics/CLI/spells/wish.md) spell or similar magic can return a bodak to its former life.

## Undead Nature

A bodak doesn't require air, food, drink, or sleep.

```statblock
"name": "Bodak (VGM)"
"size": "Medium"
"type": "undead"
"alignment": "Chaotic Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "58"
"hit_dice": "9d8 + 18"
"stats":
- !!int "15"
- !!int "16"
- !!int "15"
- !!int "7"
- !!int "12"
- !!int "12"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "4"
"damage_resistances": "cold; fire; necrotic; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "lightning, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 14"
"languages": "Abyssal, the languages it knew in life"
"cr": "6"
"traits":
- "desc": "The bodak can activate or deactivate this feature as a bonus action. While\
    \ active, the aura deals 5 necrotic damage to any creature that ends its turn\
    \ within 30 feet of the bodak. Undead and fiends ignore this effect."
  "name": "Aura of Annihilation"
- "desc": "When a creature that can see the bodak's eyes starts its turn within 30\
    \ feet of the bodak, the bodak can force it to make a DC 13 Constitution saving\
    \ throw if the bodak isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ and can see the creature. If the saving throw fails by 5 or more, the creature\
    \ is reduced to 0 hit points, unless it is immune to the [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ condition. Otherwise, a creature takes dice:3d10|text(16) (3d10) psychic\
    \ damage on a failed save.\n\nUnless [surprised](/3-Mechanics/CLI/rules/conditions.md#surprised),\
    \ a creature can avert its eyes to avoid the saving throw at the start of its\
    \ turn. If the creature does so, it has disadvantage on attack rolls against the\
    \ bodak until the start of its next turn. If the creature looks at the bodak in\
    \ the meantime, it must immediately make the saving throw."
  "name": "Death Gaze"
- "desc": "The bodak takes 5 radiant damage when it starts its turn in sunlight. While\
    \ in sunlight, it has disadvantage on attack rolls and ability checks."
  "name": "Sunlight Hypersensitivity"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage plus dice:2d8|text(9)\
    \ (2d8) necrotic damage."
  "name": "Fist"
- "desc": "One creature that the bodak can see within 60 feet of it must make a DC\
    \ 13 Constitution saving throw, taking dice:4d10|text(22) (4d10) necrotic\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Withering Gaze"
"source":
- "VGM"
- "ToA"
- "GoS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Bodak.webp"
```
^statblock

```encounter-table
name: Bodak
creatures:
 - 1: Bodak
```

## Environment

underdark, swamp, urban