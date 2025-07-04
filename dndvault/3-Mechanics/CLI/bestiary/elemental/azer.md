---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/2
- monster/size/medium
- monster/type/elemental
aliases: ["Azer"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 22. Available in the SRD.
---
# [Azer](3-Mechanics\CLI\bestiary\elemental/azer.md)
*Source: Monster Manual p. 22. Available in the SRD.*  

Natives of the Elemental Plane of Fire, azers are master crafters, expert miners, and sworn foes of the efreet. In appearance and manner, an azer resembles a male dwarf, but this is a facade. Beneath its metallic looking skin, an azer is a being of fire, which outwardly manifests in its fiery hair and beard.

## Made, Not Born

Azers don't reproduce. They are each crafted from bronze by another azer and imbued with a portion of the crafter's inner flame. Each azer is sculpted with unique features. This crafting process limits the growth of the azer population and is the primary reason that these creatures remain rare.

## Volcanic Dwellers

Azers dwell in a kingdom on the border between the Elemental Plane of Earth and the Elemental Plane of Fire-a range of mountains and volcanoes whose spires rise as a series of fortresses. Beneath mountain peaks, under volcanic calderas, and amid rivers of magma, azers extract gleaming metals and glittering gems from the earth. Squads of azer patrol the passes and tunnels of their realm, fending off the salamander raiders whose efreet masters order strikes against the azer kingdom.

## Enemies of the Efreet

Long ago, the efreet and the azers were allies. Azers helped create the City of Brass, forging that home of the efreet into one of the most wondrous places in creation. When the azers had finished their work, the efreet betrayed them, making a failed attempt to enslave the azers so as to protect the secrets of the city. Despite occasional raids and skirmishes, however, the two sides have so far refrained from all-out conflict. The azers believe that only the threat of them revealing the hidden ways into the City of Brass keeps the efreet in check.

## Masters of Metal and Gems

Azers are masterful artisans, and create beautiful works from the gems and precious metals found in their volcanic habitat. They rate the value of such treasures above all other things, sometimes dispatching parties across the planes to seek out rare metals and gemstones.

When azers are called by magic to the Material Plane, it is typically to help forge an elaborate magic item or work of art, for it is said that their skill in such craft knows no equal.

## Living Fire

An azer doesn't require food, drink, or sleep.

> [!quote]- A quote from Arakses al-Saqar, efreeti pasha  
> 
> Give me a hundred azer slaves, and I can forge an empire that would make the gods tremble.


```statblock
"name": "Azer"
"size": "Medium"
"type": "elemental"
"alignment": "Lawful Neutral"
"ac": !!int "17"
"ac_class": "natural armor, [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "39"
"hit_dice": "6d8 + 12"
"stats":
- !!int "17"
- !!int "12"
- !!int "15"
- !!int "12"
- !!int "13"
- !!int "10"
"speed": "30 ft."
"saves":
  "Constitution": !!int "4"
"damage_immunities": "fire, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "passive Perception 11"
"languages": "Ignan"
"cr": "2"
"traits":
- "desc": "A creature that touches the azer or hits it with a melee attack while within\
    \ 5 feet of it takes dice:1d10|text(5) (1d10) fire damage."
  "name": "Heated Body"
- "desc": "When the azer hits with a metal melee weapon, it deals an extra dice:1d6|text(3)\
    \ (1d6) fire damage (included in the attack)."
  "name": "Heated Weapons"
- "desc": "The azer sheds bright light in a 10-foot radius and dim light for an additional\
    \ 10 feet."
  "name": "Illumination"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) bludgeoning damage, or dice:1d10\
    \ + 3|text(8) (1d10 + 3) bludgeoning damage if used with two hands to make\
    \ a melee attack, plus dice:1d6|text(3) (1d6) fire damage."
  "name": "Warhammer"
"source":
- "MM"
- "PotA"
- "WDMM"
- "KftGV"
- "DoDk"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Azer.webp"
```
^statblock

```encounter-table
name: Azer
creatures:
 - 1: Azer
```