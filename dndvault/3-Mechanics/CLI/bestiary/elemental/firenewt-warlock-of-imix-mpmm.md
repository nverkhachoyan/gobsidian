---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/elemental
aliases: ["Firenewt Warlock of Imix"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 125, Volo's Guide to Monsters p. 143
---
# [Firenewt Warlock of Imix](3-Mechanics\CLI\bestiary\elemental/firenewt-warlock-of-imix-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 125, Volo's Guide to Monsters p. 143*  

```statblock
"name": "Firenewt Warlock of Imix (MPMM)"
"size": "Medium"
"type": "elemental"
"alignment": "Typically  Neutral Evil"
"ac": !!int "10"
"hp": !!int "33"
"hit_dice": "6d8 + 6"
"stats":
- !!int "13"
- !!int "11"
- !!int "12"
- !!int "9"
- !!int "11"
- !!int "14"
"speed": "30 ft."
"damage_immunities": "fire"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "Draconic, Ignan"
"cr": "1"
"traits":
- "desc": "The firenewt casts one of the following spells, using Charisma as the spellcasting\
    \ ability (spell save DC 12):\n\nAt will: [guidance](/3-Mechanics/CLI/spells/guidance.md),\
    \ [light](/3-Mechanics/CLI/spells/light.md), [mage armor](/3-Mechanics/CLI/spells/mage-armor.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [prestidigitation](/3-Mechanics/CLI/spells/prestidigitation.md)"
  "name": "Spellcasting"
- "desc": "The firenewt can breathe air and water."
  "name": "Amphibious"
- "desc": "Magical darkness doesn't impede the firenewt's [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision)."
  "name": "Devil's Sight"
- "desc": "When the firenewt reduces an enemy to 0 hit points, the firenewt gains\
    \ 5 temporary hit points."
  "name": "Imix's Blessing"
"actions":
- "desc": "The firenewt makes three Morningstar or Fire Ray attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) piercing damage."
  "name": "Morningstar"
- "desc": "Ranged Spell Attack: dice: d20+4 (+4) to hit, range 120 ft., one\
    \ target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) fire damage."
  "name": "Fire Ray"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Firenewt%20Warlock%20of%20Imix.webp"
```
^statblock

```encounter-table
name: Firenewt Warlock of Imix
creatures:
 - 1: Firenewt Warlock of Imix
```

## Environment

hill, mountain, underdark