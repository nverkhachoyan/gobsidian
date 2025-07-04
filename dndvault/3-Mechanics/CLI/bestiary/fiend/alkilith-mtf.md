---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/11
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/demon
aliases: ["Alkilith"]
NoteIcon: monster
BestiaryType: fiend (demon)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 130, Dragon of Icespire Peak, Sleeping Dragon's Wake
---
# [Alkilith](3-Mechanics\CLI\bestiary\fiend/alkilith-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 130, Dragon of Icespire Peak, Sleeping Dragon's Wake*  

## Alkilith

An alkilith is easily mistaken for some kind of foul fungal growth that appears on doorways, windows, and other portals. These dripping infestations conceal the demonic nature of the alkilith, making what should be a dire warning appear strange but otherwise innocuous. Wherever alkiliths take root, they weaken the fabric of reality, creating a portal through which even nastier demons can invade.

### Symptoms of Doom

The appearance of an alkilith in the world heralds a great wrongness and an imminent catastrophe. An alkilith searches for an aperture such as a window or a door around which it can take root, stretching its body around the opening and anchoring itself with a sticky secretion. If left undisturbed, the opening becomes attuned to the Abyss and eventually becomes a portal to that plane (see ""Planar Portals"" in chapter 2 of the "Dungeon Master's Guide").

### Spawn of Juiblex

Alkiliths spring from the cast-off bits of Juiblex's hideous, shuddering body, then gradually become self-aware and set out to find their way onto the Material Plane. Since most cultists consider them too risky for summoning—they can, after all, create portals to the Abyss—alkiliths must find other escape routes out of their native plane.

```statblock
"name": "Alkilith (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "demon"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "157"
"hit_dice": "15d8 + 90"
"stats":
- !!int "12"
- !!int "19"
- !!int "22"
- !!int "6"
- !!int "11"
- !!int "7"
"speed": "40 ft."
"saves":
  "Dexterity": !!int "8"
  "Constitution": !!int "10"
"skillsaves":
  "Stealth": !!int "8"
"damage_resistances": "acid; cold; fire; lightning; bludgeoning, piercing, slashing\
  \ from nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "understands Abyssal but can't speak"
"cr": "11"
"traits":
- "desc": "The alkilith can move through a space as narrow as 1 inch wide without\
    \ squeezing."
  "name": "Amorphous"
- "desc": "While the alkilith is motionless, it is indistinguishable from an ordinary\
    \ slime or fungus."
  "name": "False Appearance"
- "desc": "Any creature that isn't a demon that starts its turn within 30 feet of\
    \ the alkilith must succeed on a DC 18 Wisdom saving throw, or it hears a faint\
    \ buzzing in its head for a moment and has disadvantage on its next attack roll,\
    \ saving throw, or ability check.\n\nIf the saving throw against Foment Madness\
    \ fails by 5 or more, the creature is instead subjected to the [confusion](/3-Mechanics/CLI/spells/confusion.md)\
    \ spell for 1 minute (no [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ required by the alkilith). While under the effect of that confusion, the creature\
    \ is immune to Foment Madness."
  "name": "Foment Madness"
- "desc": "The alkilith has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The alkilith makes three tentacle attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 15 ft., one target.\
    \ Hit: dice:4d6 + 4|text(18) (4d6 + 4) acid damage."
  "name": "Tentacle"
"source":
- "MTF"
- "DIP"
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Alkilith.webp"
```
^statblock

```encounter-table
name: Alkilith
creatures:
 - 1: Alkilith
```

## Environment

underdark, urban