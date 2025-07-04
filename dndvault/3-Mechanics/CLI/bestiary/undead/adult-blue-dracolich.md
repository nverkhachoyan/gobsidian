---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/17
- monster/environment/desert
- monster/size/huge
- monster/type/undead
aliases: ["Adult Blue Dracolich"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 84
---
# [Adult Blue Dracolich](3-Mechanics\CLI\bestiary\undead/adult-blue-dracolich.md)
*Source: Monster Manual p. 84*  

Even as long-lived as they are, all dragons must eventually die. This thought doesn't sit well with many dragons, some of which allow themselves to be transformed by necromantic energy and ancient rituals into powerful undead dracoliches. Only the most narcissistic dragons choose this path, knowing that by doing so, they sever all ties to their kin and the dragon gods.

## Beyond Death

A dracolich retains its shape and size upon transforming, its skin and scales drawing tight to its bones or sloughing away to leave a skeletal form behind. Its eyes appear as glowing points of light floating in shadowy sockets, hinting at the malevolence of its undead mind.

Though many dragons pursue vain goals of destruction and dominance, dracoliches are more nefarious than the most evil dragons, driven to rule over all. A dracolich is a fiendishly intelligent tyrant that crafts complex webs of foul schemes, attracting servants motivated by greed and a lust for power. Acting from the shadows and actively plotting to keep its existence a secret, a dracolich is a cunning and challenging foe.

## Dracolich Phylacteries

Creating a dracolich requires the cooperation of the dragon and a group of mages or cultists that can perform the proper ritual. During the ritual, the dragon consumes a toxic brew that slays it instantly. The attendant spellcasters then ensnare its spirit and transfer it to a special gemstone that functions like a lich's phylactery. As the dragon's flesh rots away, the spirit inside the gem returns to animate the dragon's bones.

If a dracolich's physical form is ever destroyed, its spirit returns to the gem as long as the two are on the same plane. If the gem comes into contact with another dragon's corpse, the dracolich's spirit can take possession of that corpse to become a new dracolich. If the dracolich's spirit gem is taken to another plane, the dracolich's spirit has nowhere to go when its undead body is destroyed and simply passes into the afterlife.

## Dracolich Template

Only an ancient or adult true dragon can be transformed into a dracolich. Younger dragons that attempt to undergo the transformation die, as do other creatures that aren't true dragons but possess the dragon type, such as pseudodragons and wyverns. A shadow dragon can't be transformed into a dracolich, for it has already lost too much of its physical form.

When a dragon becomes a dracolich, it retains its statistics except as described below. The dragon loses any trait, such as Amphibious, that assumes a living physiology. The dracolich might retain or lose any or all of its lair actions or inherit new ones, as the DM sees fit.

### Type

The dracolich's type changes from dragon to undead, and it no longer requires air, food, drink, or sleep.

### Damage Resistance

The dracolich has resistance to necrotic damage.

### Damage Immunities

The dracolich has immunity to poison. It also retains any immunities it had prior to becoming a dracolich.

### Condition Immunities

The dracolich can't be [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), or [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). It also doesn't suffer from [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion).

### Magic Resistance

The dracolich has advantage on saving throws against spells and other magical effects.

```statblock
"name": "Adult Blue Dracolich"
"size": "Huge"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "225"
"hit_dice": "18d12 + 108"
"stats":
- !!int "25"
- !!int "10"
- !!int "23"
- !!int "16"
- !!int "15"
- !!int "19"
"speed": "40 ft., burrow 30 ft., fly 80 ft."
"saves":
  "Charisma": !!int "10"
  "Dexterity": !!int "6"
  "Wisdom": !!int "8"
  "Constitution": !!int "12"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "14"
"damage_resistances": "necrotic"
"damage_immunities": "lightning, poison"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion), [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 120 ft., passive Perception 24"
"languages": "Common, Draconic"
"cr": "17"
"traits":
- "desc": "If the dracolich fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The dracolich has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
"actions":
- "desc": "The dracolich can use its Frightful Presence. It then makes three attacks:\
    \ one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d10 + 7|text(18) (2d10 + 7) piercing damage plus dice:1d10|text(5)\
    \ (1d10) lightning damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 7|text(14) (2d6 + 7) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+13 (+13) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d8 + 7|text(16) (2d8 + 7) bludgeoning damage."
  "name": "Tail"
- "desc": "Each creature of the dracolich's choice that is within 120 feet of the\
    \ dracolich and aware of it must succeed on a DC 18 Wisdom saving throw or become\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) for 1 minute.\
    \ A creature can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success. If a creature's saving throw is successful\
    \ or the effect ends for it, the creature is immune to the dracolich's Frightful\
    \ Presence for the next 24 hours."
  "name": "Frightful Presence"
- "desc": "The dracolich exhales lightning in a 90-foot line that is 5 feet wide.\
    \ Each creature in that line must make a DC 20 Dexterity saving throw, taking\
    \ dice:12d10|text(66) (12d10) lightning damage on a failed save, or half as\
    \ much damage on a successful one."
  "name": "Lightning Breath (Recharge 5-6)"
"legendary_actions":
- "desc": "The dracolich makes a Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ check."
  "name": "Detect"
- "desc": "The dracolich makes a tail attack."
  "name": "Tail Attack"
- "desc": "The dracolich beats its tattered wings. Each creature within 10 feet of\
    \ the dracolich must succeed on a DC 21 Dexterity saving throw or take dice:2d6\
    \ + 7|text(14) (2d6 + 7) bludgeoning damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ After beating its wings this way, the dracolich can fly up to half its flying\
    \ speed."
  "name": "Wing Attack (Costs 2 Actions)"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), the dragon takes a lair\
    \ action to cause one of the following effects; the dragon can't use the same\
    \ effect two rounds in a row:"
  "name": ""
- "desc": "- Part of the ceiling collapses above one creature that the dragon can\
    \ see within 120 feet of it. The creature must succeed on a DC 15 Dexterity saving\
    \ throw or take dice:3d6|text(10) (3d6) bludgeoning damage and be knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone) and buried. The buried target\
    \ is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) and unable\
    \ to breathe or stand up. A creature can take an action to make a DC 10 Strength\
    \ check, ending the buried state on a success.  \n- A cloud of sand swirls about\
    \ in a 20-foot-radius sphere centered on a point the dragon can see within 120\
    \ feet of it. The cloud spreads around corners. Each creature in the cloud must\
    \ succeed on a DC 15 Constitution saving throw or be [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ for 1 minute. A creature can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success.  \n- Lightning arcs, forming\
    \ a 5-foot-wide line between two of the lair's solid surfaces that the dragon\
    \ can see. They must be within 120 feet of the dragon and 120 feet of each other.\
    \ Each creature in that line must succeed on a DC 15 Dexterity saving throw or\
    \ take dice:3d6|text(10) (3d6) lightning damage.  "
  "name": ""
- "desc": "At your discretion, a legendary ([adult](/3-Mechanics/CLI/bestiary/dragon/adult-blue-dragon.md)\
    \ or [ancient](/3-Mechanics/CLI/bestiary/dragon/ancient-blue-dragon.md)) blue\
    \ dragon can use one or both of the following additional lair actions while in\
    \ its lair:\n\n- Misleading Mirage. Until initiative count 20 on the next\
    \ round, at each intersection or branching passage in the lair, a creature other\
    \ than the dragon has a 50% chance chance of going in a different direction from\
    \ the direction it intends.  \n- Sudden Sinkhole. The dragon chooses a point\
    \ on the ground that it can see within 120 feet of it. A 5-foot-radius, 20-foot-deep\
    \ pit forms centered on that point. A creature on the ground above where the pit\
    \ formed must make a DC 15 Dexterity saving throw. On a failed save, a creature\
    \ falls to the bottom of the pit, taking dice:2d6|text(7) (2d6) bludgeoning\
    \ damage and landing [prone](/3-Mechanics/CLI/rules/conditions.md#prone). On a\
    \ successful save, a creature moves to the nearest unoccupied space instead of\
    \ falling in the pit.  "
  "name": "Additional Lair Actions"
"regional_effects":
- "desc": "The region containing a legendary blue dragon's lair is warped by the dragon's\
    \ magic, which creates one or more of the following effects:"
  "name": ""
- "desc": "- Thunderstorms rage within 6 miles of the lair.  \n- Dust devils scour\
    \ the land within 6 miles of the lair. A dust devil has the statistics of an air\
    \ elemental, but it can't fly, has a speed of 50 feet, and has an Intelligence\
    \ and Charisma of 1 (-5).  \n- Hidden sinkholes form in and around the dragon's\
    \ lair. A sinkhole can be spotted from a safe distance with a successful DC 20\
    \ Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) check. Otherwise,\
    \ the first creature to step on the thin crust covering the sinkhole must succeed\
    \ on a DC 15 Dexterity saving throw or fall 1d6 × 10 feet into the sinkhole.\
    \  "
  "name": ""
- "desc": "If the dragon dies, the dust devils disappear immediately, and the thunderstorms\
    \ abate within dice: 1d10|avg|noform (1d10) days. Any sinkholes remain where\
    \ they are."
  "name": ""
- "desc": "Any of these effects might appear in the area around a blue dragon's lair,\
    \ in addition to or instead of the effects described in the Monster Manual:\n\
    \n- Blue Luster. Creatures that spend a year within 1 mile of the dragon's\
    \ lair find blue objects fascinating and feel compelled to acquire them at every\
    \ opportunity.  \n- Mirage Terrain. The area immediately surrounding the lair\
    \ appears to be a lush oasis. A creature carefully examining the illusion can\
    \ attempt a DC 15 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check to disbelieve it. A creature who disbelieves the illusion sees it as a\
    \ vague image superimposed on the underlying terrain.  \n- Sandstorm. A sandstorm\
    \ blows constantly within 1 mile of the dragon's lair.  "
  "name": "Additional Regional Effects"
"source":
- "MM"
- "WDMM"
- "CM"
- "JttRC"
- "SatO"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Adult%20Blue%20Dracolich.webp"
```
^statblock

```encounter-table
name: Adult Blue Dracolich
creatures:
 - 1: Adult Blue Dracolich
```

## Environment

desert