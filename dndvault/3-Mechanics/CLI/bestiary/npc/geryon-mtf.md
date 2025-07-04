---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/22
- monster/size/huge
- monster/type/fiend/devil
aliases: ["Geryon"]
NoteIcon: monster
BestiaryType: fiend (devil)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 173
---
# [Geryon](3-Mechanics\CLI\bestiary\npc/geryon-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 173*  

> [!quote]- A quote from Mordenkainen  
> 
> Which is less worthy: the archdevil who leads a layer while being trapped in a block of ice, or the archdevil who can't outmaneuver a frozen adversary?

## Geryon

Geryon is locked in an endless struggle with Levistus for control of Stygia. The two have fought each other for centuries, each displacing the other innumerable times. Currently, Geryon occupies an odd position in the infernal hierarchy. Although Levistus still claims lordship over Stygia, he has been trapped in an enormous block of ice at the command of Asmodeus. For his part, Geryon marshals his followers and seeks to discover the means to replace his hated rival.

Among the archdevils, Geryon and Zariel are especially known for martial prowess. He is a ferocious hunter and a relentless tracker. Other devils command legions and bid their followers to battle their enemies. Geryon loves the feeling of flesh and steel being sundered beneath his claws, and the taste of his foes' blood.

His ferocity serves him well in Stygia's frozen waste, but it has also limited his ability to collect souls and forge an effective hierarchy. Sages who study the Nine Hells believe that the battle for control of Stygia is a test staged by Asmodeus in hopes of purging the worst impulses from both Geryon and Levistus, or at the very least opening the door for a competent replacement for both to rise from the ranks.

## Geryon's Lair

Geryon has recently reclaimed his ancient fortress, Coldsteel, a sprawling complex that rises from the ice and snow at the center of Stygia. He roams the passages of this place, scattering the [ice devils](/3-Mechanics/CLI/bestiary/fiend/ice-devil.md) and [minotaur](/3-Mechanics/CLI/bestiary/monstrosity/minotaur.md) slaves he took from [Baphomet](/3-Mechanics/CLI/bestiary/npc/baphomet-mtf.md), raging against Asmodeus's betrayal while spitting oaths of vengeance and hatching mad schemes to reclaim his standing from Levistus.

```statblock
"name": "Geryon (MTF)"
"size": "Huge"
"type": "fiend"
"subtype": "devil"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "300"
"hit_dice": "24d12 + 144"
"stats":
- !!int "29"
- !!int "17"
- !!int "22"
- !!int "19"
- !!int "16"
- !!int "23"
"speed": "30 ft., fly 50 ft."
"saves":
  "Charisma": !!int "13"
  "Dexterity": !!int "10"
  "Wisdom": !!int "10"
  "Constitution": !!int "13"
"skillsaves":
  "Intimidation": !!int "13"
  "Deception": !!int "13"
  "Perception": !!int "10"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks that\
  \ aren't silvered"
"damage_immunities": "cold, fire, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "truesight 120 ft., passive Perception 20"
"languages": "all, telepathy 120 ft."
"cr": "22"
"traits":
- "desc": "Geryon's innate spellcasting ability is Charisma (spell save DC 21). He\
    \ can innately cast the following spells, requiring no material components:\n\n\
    At will: [alter self](/3-Mechanics/CLI/spells/alter-self.md) (can become Medium\
    \ when changing his appearance), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [geas](/3-Mechanics/CLI/spells/geas.md), [ice storm](/3-Mechanics/CLI/spells/ice-storm.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md) (self only), [locate\
    \ object](/3-Mechanics/CLI/spells/locate-object.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md),\
    \ [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)\n\n1/day each: [divine\
    \ word](/3-Mechanics/CLI/spells/divine-word.md), [symbol](/3-Mechanics/CLI/spells/symbol.md)\
    \ (pain only)"
  "name": "Innate Spellcasting"
- "desc": "If Geryon fails a saving throw, he can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Geryon has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "Geryon's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "Geryon regains 20 hit points at the start of his turn. If he takes radiant\
    \ damage, this trait doesn't function at the start of his next turn. Geryon dies\
    \ only if he starts his turn with 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "Geryon makes two attacks: one with his claws and one with his stinger."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 15 ft., one\
    \ target. Hit: dice:4d6 + 9|text(23) (4d6 + 9) slashing damage. If the target\
    \ is Large or smaller, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (DC 24) and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the grapple ends. Geryon can grapple one creature at a time. If the target\
    \ is already [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) by Geryon,\
    \ the target takes an extra dice:6d8|text(27) (6d8) slashing damage."
  "name": "Claws"
- "desc": "Melee Weapon Attack: dice: d20+16 (+16) to hit, reach 20 ft., one\
    \ creature. Hit: dice:2d4 + 9|text(14) (2d4 + 9) piercing damage, and the\
    \ target must succeed on a DC 21 Constitution saving throw or take dice:2d12|text(13)\
    \ (2d12) poison damage and become [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ until it finishes a short or long rest. The target's hit point maximum is reduced\
    \ by an amount equal to half the poison damage it takes. If its hit point maximum\
    \ drops to 0, it dies. This reduction lasts until the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ condition is removed."
  "name": "Stinger"
- "desc": "Geryon magically teleports, along with any equipment he is wearing and\
    \ carrying, up to 120 feet to an unoccupied space he can see."
  "name": "Teleport"
"legendary_actions":
- "desc": "Geryon targets one creature he can see within 60 feet of him. If the target\
    \ can see Geryon, the target must succeed on a DC 23 Wisdom saving throw or become\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) of Geryon until\
    \ the end of its next turn."
  "name": "Infernal Glare"
- "desc": "Geryon attacks with his stinger."
  "name": "Swift Sting (Costs 2 Actions)"
- "desc": "Geryon uses his Teleport action."
  "name": "Teleport"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), Geryon can take a lair\
    \ action to cause one of the following effects; he can't use the same effect two\
    \ rounds in a row:"
  "name": ""
- "desc": "- Geryon causes a blast of cold to burst from the ground at a point he\
    \ can see within 120 feet of him. The cold fills a cube, 10 feet on each side,\
    \ centered on that point. Each creature in that area must succeed on a DC 21 Constitution\
    \ saving throw or take dice:8d6|text(28) (8d6) cold damage.  \n- Geryon targets\
    \ one creature he can see within 60 feet of him. The target must succeed on a\
    \ DC 21 Wisdom saving throw or become [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ for 1 minute. The target can end the effect on itself if it deals any damage\
    \ to one or more of its allies.  \n- Geryon casts the [banishment](/3-Mechanics/CLI/spells/banishment.md)\
    \ spell.  "
  "name": ""
"regional_effects":
- "desc": "The region containing Geryon's lair is warped by his magic, creating one\
    \ or more of the following effects:"
  "name": ""
- "desc": "- Intelligent creatures within 1 mile of the lair frequently see shimmering\
    \ portals leading to places they consider safe. Passing through a portal always\
    \ deposits a traveler somewhere in Stygia.  \n- Freezing strong winds howl around\
    \ the area within 1 mile of the lair.  \n- Howls and screams fill the air within\
    \ 1 mile of the lair. Any creature that finishes a short or long rest in this\
    \ area must succeed on a DC 21 Wisdom saving throw or derive no benefit from the\
    \ rest.  "
  "name": ""
- "desc": "If Geryon dies, these effects fade over the course of dice: 1d10|avg|noform\
    \ (1d10) days."
  "name": ""
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Geryon.webp"
```
^statblock

```encounter-table
name: Geryon
creatures:
 - 1: Geryon
```