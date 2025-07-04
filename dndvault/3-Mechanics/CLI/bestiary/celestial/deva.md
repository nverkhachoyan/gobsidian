---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/10
- monster/size/medium
- monster/type/celestial
aliases: ["Deva"]
NoteIcon: monster
BestiaryType: celestial
SourceType: Bestiary
BookSource: Monster Manual p. 16. Available in the SRD.
---
# [Deva](3-Mechanics\CLI\bestiary\celestial/deva.md)
*Source: Monster Manual p. 16. Available in the SRD.*  

An angel is a celestial agent sent forth into the planes to further its god's agenda for weal or woe. Its sublime beauty and presence can drive awestruck onlookers to their knees. Yet angels are destroyers too, and their appearance portends doom as often as it signals hope.

## Shards of the Divine

Angels are formed from the astral essence of benevolent gods and are thus divine beings of great power and foresight.

Angels act out the will of their gods with tireless devotion. Even chaotic good deities command lawful good angels, knowing that the angels' dedication to order best allows them to fulfill divine commands. An angel follows a single driving purpose, as decreed by its deity. However, an angel is incapable of following commands that stray from the path of law and good.

An angel slays evil creatures without remorse. As the embodiment of law and good, an angel is almost never mistaken in its judgments. This quality can create a sense of superiority in an angel, a sense that comes to the fore when an angel's task conflicts with the goals of another creature. The angel never acquiesces or gives way. When an angel is sent to aid mortals, it is sent not to serve but to command. The gods of good therefore send their angels among mortals only in response to the most dire circumstances.

## Fallen Angels

An angel's moral compass grants it a sense of infallibility that can sometimes spell its undoing. Angels are usually too wise to fall for a simple deception, but sometimes pride can lead one to commit an evil act. Whether intentional or accidental, such an act is a permanent stain that marks the angel as an outcast. Fallen angels retain their power but lose their connection to the deities from which they were made. Most fallen angels take their banishment personally, rebelling against the powers they served by seeking rulership over a section of the Abyss or a place among other fallen in the hierarchy of the Nine Hells. Zariel, the ruler of the first layer of the Nine Hells, is such a creature. Rather than rebel, some fallen angels resign themselves to an isolated existence on the Material Plane, living in disguise as simple hermits. If they are redeemed, they can become powerful allies dedicated to justice and compassionate service.

## Immortal Nature

An angel doesn't require food, drink, or sleep.

## Deva

Devas are angels that act as divine messengers or agents to the Material Plane, the Shadowfell, and the Feywild and that can assume a form appropriate to the realm they are sent to.

Legend tells of angels that take mortal form for years, lending aid, hope, and courage to goodhearted folk. A deva can take any shape, although it prefers to appear to mortals as an innocuous humanoid or animal. When circumstances require that it cast off its guise, a deva is a beautiful humanoid-like creature with silvery skin. Its hair and eyes gleam with an unearthly luster, and large feathery wings unfurl from its shoulder blades.

```statblock
"name": "Deva"
"size": "Medium"
"type": "celestial"
"alignment": "Lawful Good"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "136"
"hit_dice": "16d8 + 64"
"stats":
- !!int "18"
- !!int "18"
- !!int "18"
- !!int "17"
- !!int "20"
- !!int "20"
"speed": "30 ft., fly 90 ft."
"saves":
  "Charisma": !!int "9"
  "Wisdom": !!int "9"
"skillsaves":
  "Insight": !!int "9"
  "Perception": !!int "9"
"damage_resistances": "radiant; bludgeoning, piercing, slashing from nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "darkvision 120 ft., passive Perception 19"
"languages": "all, telepathy 120 ft."
"cr": "10"
"traits":
- "desc": "The deva's spellcasting ability is Charisma (spell save DC 17). The deva\
    \ can innately cast the following spells, requiring only verbal components:\n\n\
    At will: [detect evil and good](/3-Mechanics/CLI/spells/detect-evil-and-good.md)\n\
    \n1/day each: [commune](/3-Mechanics/CLI/spells/commune.md), [raise dead](/3-Mechanics/CLI/spells/raise-dead.md)"
  "name": "Innate Spellcasting"
- "desc": "The deva's weapon attacks are magical. When the deva hits with any weapon,\
    \ the weapon deals an extra dice: 4d8|avg|noform (4d8) radiant damage (included\
    \ in the attack)."
  "name": "Angelic Weapons"
- "desc": "The deva has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The deva makes two melee attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) bludgeoning damage plus dice:4d8|text(18)\
    \ (4d8) radiant damage."
  "name": "Mace"
- "desc": "The deva touches another creature. The target magically regains dice:4d8\
    \ + 2|text(20) (4d8 + 2) hit points and is freed from any curse, disease, poison,\
    \ blindness, or deafness."
  "name": "Healing Touch (3/Day)"
- "desc": "The deva magically polymorphs into a humanoid or beast that has a challenge\
    \ rating equal to or less than its own, or back into its true form. It reverts\
    \ to its true form if it dies. Any equipment it is wearing or carrying is absorbed\
    \ or borne by the new form (the deva's choice).\n\nIn a new form, the deva retains\
    \ its game statistics and ability to speak, but its AC, movement modes, Strength,\
    \ Dexterity, and special senses are replaced by those of the new form, and it\
    \ gains any statistics and capabilities (except class features, legendary actions,\
    \ and lair actions) that the new form has but that it lacks."
  "name": "Change Shape"
"source":
- "MM"
- "WDMM"
- "CoS"
- "CRCotN"
- "JttRC"
- "KftGV"
- "SatO"
- "ToFW"
- "BMT"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Deva.webp"
```
^statblock

```encounter-table
name: Deva
creatures:
 - 1: Deva
```