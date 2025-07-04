---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/sdw
- monster/cr/16
- monster/size/huge
- monster/type/dragon
aliases: ["Lhammaruntosz"]
NoteIcon: monster
BestiaryType: dragon
SourceType: Bestiary
BookSource: Sleeping Dragon's Wake
---
# [Lhammaruntosz](3-Mechanics\CLI\bestiary\npc/lhammaruntosz-sdw.md)
*Source: Sleeping Dragon's Wake*  

Lhammaruntosz, a bronze dragon with the ability to heal quickly, spent decades defending Leilon and the surrounding area as the captain of the Scaly Eye, a fleet that battled pirates and other threats. To honor her deeds, the Swords of Leilon constructed the Bronze Shrine, a massive temple to Bahamut, god of metallic dragons, in a cliff overlooking the sea. The shrine's face is carved in Lhammaruntosz's likeness and includes quarters for the rest of the Scaly Eye and a magic statue of Bahamut, which the dragon can use to commune with the deity.

In recent decades Lhammaruntosz has retreated inside the shrine, becoming reclusive due to a attack by a disguised demon which has driven her mad. She leaves on rare occasions to hunt for food, returning as soon as possible. Members of the Scaly Eye still live within the Bronze Shrine, as Lhammaruntosz has ordered them to stay on as her guardians.

```statblock
"name": "Lhammaruntosz (SDW)"
"size": "Huge"
"type": "dragon"
"alignment": "Lawful Good"
"ac": !!int "19"
"ac_class": "natural armor"
"hp": !!int "212"
"hit_dice": "17d12 + 102"
"stats":
- !!int "25"
- !!int "10"
- !!int "23"
- !!int "16"
- !!int "15"
- !!int "19"
"speed": "40 ft., fly 80 ft., swim 40 ft."
"saves":
  "Charisma": !!int "9"
  "Dexterity": !!int "5"
  "Wisdom": !!int "7"
  "Constitution": !!int "11"
"skillsaves":
  "Stealth": !!int "5"
  "Insight": !!int "7"
  "Perception": !!int "12"
"damage_immunities": "lightning"
"senses": "blindsight 60 ft., darkvision 120 ft., passive Perception 22"
"languages": "Common, Draconic"
"cr": "16"
"traits":
- "desc": "Lhammaruntosz's spellcasting ability is Charisma (spell save DC 17). She\
    \ can innately cast the following spells, requiring no material components:\n\n\
    1/day each: [create food and water](/3-Mechanics/CLI/spells/create-food-and-water.md),\
    \ [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md),\
    \ [speak with animals](/3-Mechanics/CLI/spells/speak-with-animals.md)"
  "name": "Innate Spellcasting"
- "desc": "Lhammaruntosz is an 8th-level spellcaster. Her spellcasting ability is\
    \ Charisma (spell save DC 17; dice: d20+9 (+9) to hit with spell attacks).\
    \ She has the following sorcerer spells prepared:\n\nCantrips (at will): [light](/3-Mechanics/CLI/spells/light.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [mending](/3-Mechanics/CLI/spells/mending.md)\n\
    \n1st level (4 slots): [charm person](/3-Mechanics/CLI/spells/charm-person.md),\
    \ [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [expeditious retreat](/3-Mechanics/CLI/spells/expeditious-retreat.md),\
    \ [sleep](/3-Mechanics/CLI/spells/sleep.md)\n\n2nd level (3 slots): [darkness](/3-Mechanics/CLI/spells/darkness.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [suggestion](/3-Mechanics/CLI/spells/suggestion.md)\n\
    \n3rd level (3 slots): [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ [protection from energy](/3-Mechanics/CLI/spells/protection-from-energy.md)\n\
    \n4th level (2 slots): [dimension door](/3-Mechanics/CLI/spells/dimension-door.md),\
    \ [stoneskin](/3-Mechanics/CLI/spells/stoneskin.md)"
  "name": "Spellcasting"
- "desc": "Lhammaruntosz can breathe air and water."
  "name": "Amphibious"
- "desc": "If Lhammaruntosz fails a saving throw, she can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "Lhammaruntosz regains 5 hit points at the start of her turn."
  "name": "Regeneration"
"actions":
- "desc": "Lhammaruntosz can use her Frightful Presence. She then makes three attacks:\
    \ one with her bite and two with her claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 10 ft., one\
    \ target. Hit: dice:2d10 + 7|text(18) (2d10 + 7) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 7|text(14) (2d6 + 7) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+12 (+12) to hit, reach 15 ft., one\
    \ target. Hit: dice:2d8 + 7|text(16) (2d8 + 7) bludgeoning damage."
  "name": "Tail"
- "desc": "Each creature of Lhammaruntosz's choice that is within 120 feet of her\
    \ and aware of her must succeed on a DC 17 Wisdom saving throw or become [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ for 1 minute. A creature can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success. If a creature's saving throw\
    \ is successful or the effect ends for it, the creature is immune to Lhammaruntosz's\
    \ Frightful Presence for the next 24 hours."
  "name": "Frightful Presence"
- "desc": "Lhammaruntosz uses one of the following breath weapons."
  "name": "Breath Weapons (Recharge 5-6)"
- "desc": "Lhammaruntosz exhales lightning in a 90-foot line that is 5 feet wide.\
    \ Each creature in that line must make a DC 19 Dexterity saving throw, taking\
    \ dice:12d10|text(66) (12d10) lightning damage on a failed save, or half as\
    \ much damage on a successful one."
  "name": "Lightning Breath"
- "desc": "Lhammaruntosz exhales repulsion energy in a 30-foot cone. Each creature\
    \ in that area must succeed on a DC 19 Strength saving throw. On a failed save,\
    \ the creature is pushed 60 feet away from the dragon."
  "name": "Repulsion Breath"
- "desc": "Lhammaruntosz magically polymorphs into a humanoid or beast that has a\
    \ challenge rating no higher than her own, or back into her true form. She reverts\
    \ to her true form if she dies. Any equipment she is wearing or carrying is absorbed\
    \ or borne by the new form (Lhammaruntosz's choice).\n\nIn a new form, Lhammaruntosz\
    \ retains her alignment, hit points, Hit Dice, ability to speak, proficiencies,\
    \ Legendary Resistance, lair actions, and Intelligence, Wisdom, and Charisma scores,\
    \ as well as this action. Her statistics and capabilities are otherwise replaced\
    \ by those of the new form, except any class features or legendary actions of\
    \ that form."
  "name": "Change Shape"
"legendary_actions":
- "desc": "Lhammaruntosz makes a Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ check."
  "name": "Detect"
- "desc": "Lhammaruntosz makes a tail attack."
  "name": "Tail Attack"
- "desc": "Lhammaruntosz beats its wings. Each creature within 10 feet of Lhammaruntosz\
    \ must succeed on a DC 20 Dexterity saving throw or take dice:2d6 + 7|text(14)\
    \ (2d6 + 7) bludgeoning damage and be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ Lhammaruntosz can then fly up to half her flying speed."
  "name": "Wing Attack (Costs 2 Actions)"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), the dragon takes a lair\
    \ action to cause one of the following effects:"
  "name": ""
- "desc": "- The dragon creates fog as though it had cast the [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md)\
    \ spell. The fog lasts until initiative count 20 on the next round.  \n- A thunderclap\
    \ originates at a point the dragon can see within 120 feet of it. Each creature\
    \ within a 20-foot radius centered on that point must make a DC 15 Constitution\
    \ saving throw or take dice:1d10|text(5) (1d10) thunder damage and be [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened)\
    \ until the end of its next turn.  "
  "name": ""
- "desc": "At your discretion, a legendary ([adult](/3-Mechanics/CLI/bestiary/dragon/adult-bronze-dragon.md)\
    \ or [ancient](/3-Mechanics/CLI/bestiary/dragon/ancient-bronze-dragon.md)) bronze\
    \ dragon can use one or more of the following additional lair actions while in\
    \ its lair:\n\n- Ocean's Call. The dragon conjures a swarm of spectral dolphins.\
    \ Each creature in the water within 120 feet of the dragon must succeed on a DC\
    \ 15 Dexterity saving throw or take dice:1d10|text(5) (1d10) slashing damage;\
    \ then the swarm vanishes.  \n- Salt Burst. The dragon chooses a point it\
    \ can see in the lair. The air in a 20-foot-radius sphere centered on that point\
    \ bursts with abrasive salt crystals. Each creature in that area must succeed\
    \ on a DC 15 Dexterity saving throw or take dice:3d6|text(10) (3d6) slashing\
    \ damage.  \n- Whelming Water. The dragon causes a strong current to course\
    \ through the water in its lair. The dragon chooses any number of creatures it\
    \ can see that are standing or swimming in water within 120 feet of it. Each chosen\
    \ creature must succeed on a DC 15 Strength saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \  "
  "name": "Additional Lair Actions"
"regional_effects":
- "desc": "The region containing a legendary bronze dragon's lair is warped by the\
    \ dragon's magic."
  "name": ""
- "desc": "- Once per day, the dragon can alter the weather in a 6-mile radius centered\
    \ on its lair. The dragon doesn't need to be outdoors; otherwise the effect is\
    \ identical to the [control weather](/3-Mechanics/CLI/spells/control-weather.md)\
    \ spell.  \n- Underwater plants within 6 miles of the dragon's lair take on dazzlingly\
    \ brilliant hues.  \n- Within its lair, the dragon can set illusory sounds, such\
    \ as soft music and strange echoes, so that they can be heard in various parts\
    \ of the lair.  "
  "name": ""
- "desc": "If the dragon dies, changed weather reverts to normal, as described in\
    \ the spell, and the other effects fade in dice: 1d10|avg|noform (1d10) days."
  "name": ""
- "desc": "Any of these effects might appear in the area around a bronze dragon's\
    \ lair, in addition to or instead of the effects described in the Monster Manual:\n\
    \n- Phantom Escort. Ghostly naval ships from an ancient armada appear, escorting\
    \ well-meaning creatures in need of the dragon's help toward the dragon's lair.\
    \  \n- Underwater Pursuit. Sailors glimpse the shadowy, illusory form of a\
    \ dragon in the depths below them, keeping pace with their vessel.  \n- Unfailing\
    \ Faithfulness. Sapient creatures that spend a year within 10 miles of the dragon's\
    \ lair find it nearly impossible to break a promise.  "
  "name": "Additional Regional Effects"
"source":
- "SDW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/SDW/Lhammaruntosz.webp"
```
^statblock

```encounter-table
name: Lhammaruntosz
creatures:
 - 1: Lhammaruntosz
```