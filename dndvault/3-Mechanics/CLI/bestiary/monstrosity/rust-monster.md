---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/environment/underdark
- monster/size/medium
- monster/type/monstrosity
aliases: ["Rust Monster"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Monster Manual p. 262. Available in the SRD.
---
# [Rust Monster](3-Mechanics\CLI\bestiary\monstrosity/rust-monster.md)
*Source: Monster Manual p. 262. Available in the SRD.*  

Most dwarves would rather face a squad of orcs than confront a single rust monster. These strange, normally docile creatures corrode ferrous metals, then gobble up the rust they create. In doing so, they have ruined the armor, shields, and weapons of countless adventurers.

A rust monster's body is covered in thick, lumpy armor, its long tail ends in a bony protrusion, and two feathery antennae sprout from its insectile head.

## Underground Scavengers

Rust monsters roam subterranean passages in search of ferrous metals such as iron, steel, adamantine, and mithral to consume.

They ignore creatures not carrying such metals, but can become aggressive toward those bearing steel weapons and armor. A rust monster can smell its food at a distance, immediately dashing toward the scent's source to corrode and consume the object.

A rust monster doesn't care if the rust it consumes comes from a spike or a sword. Adventurers can distract the creature by dropping ferrous objects behind them.

## Subterranean Wanderers

Rust monsters are rarely found in large numbers, preferring to hunt alone or in small groups. They meander along tunnels, moving from cave to cave in their tireless search for ferrous metals to consume. Their wanderings often bring them into contact with other Underdark denizens that find them harmless or unappetizing. Thus, rust monsters may be found in close proximity to other subterranean monsters. If they are well treated and well fed, they can also become friendly companions or pets.

> [!quote]- A quote from Brawn Thunderstones, dwarf paladin of Dumathoin  
> 
> Destroyed me hammer and me axe, it did. Put me faith to the test that day, it did.


```statblock
"name": "Rust Monster"
"size": "Medium"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "27"
"hit_dice": "5d8 + 5"
"stats":
- !!int "13"
- !!int "12"
- !!int "13"
- !!int "2"
- !!int "13"
- !!int "6"
"speed": "40 ft."
"senses": "darkvision 60 ft., passive Perception 11"
"languages": ""
"cr": "1/2"
"traits":
- "desc": "The rust monster can pinpoint, by scent, the location of ferrous metal\
    \ within 30 feet of it."
  "name": "Iron Scent"
- "desc": "Any nonmagical weapon made of metal that hits the rust monster corrodes.\
    \ After dealing damage, the weapon takes a permanent and cumulative −1 penalty\
    \ to damage rolls. If its penalty drops to −5, the weapon is destroyed. Non magical\
    \ ammunition made of metal that hits the rust monster is destroyed after dealing\
    \ damage."
  "name": "Rust Metal"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 1|text(5) (1d8 + 1) piercing damage."
  "name": "Bite"
- "desc": "The rust monster corrodes a nonmagical ferrous metal object it can see\
    \ within 5 feet of it. If the object isn't being worn or carried, the touch destroys\
    \ a 1-foot cube of it. If the object is being worn or carried by a creature, the\
    \ creature can make a DC 11 Dexterity saving throw to avoid the rust monster's\
    \ touch.\n\nIf the object touched is either metal armor or a metal shield being\
    \ worn or carried, it takes a permanent and cumulative −1 penalty to the AC it\
    \ offers. Armor reduced to an AC of 10 or a shield that drops to a +0 bonus is\
    \ destroyed. If the object touched is a held metal weapon, it rusts as described\
    \ in the Rust Metal trait."
  "name": "Antennae"
"source":
- "MM"
- "PotA"
- "SKT"
- "WDH"
- "WDMM"
- "IMR"
- "EGW"
- "IDRotF"
- "SatO"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Rust%20Monster.webp"
```
^statblock

```encounter-table
name: Rust Monster
creatures:
 - 1: Rust Monster
```

## Environment

underdark