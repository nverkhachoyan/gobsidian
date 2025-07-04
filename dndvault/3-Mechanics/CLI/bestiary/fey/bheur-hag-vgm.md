---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/7
- monster/environment/arctic
- monster/size/medium
- monster/type/fey
aliases: ["Bheur Hag"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 160
---
# [Bheur Hag](3-Mechanics\CLI\bestiary\fey/bheur-hag-vgm.md)
*Source: Volo's Guide to Monsters p. 160*  

Bheur hags live in wintry lands, favoring snow-covered mountains. They become more active during winter, using their ice and weather magic to make life miserable for nearby settlements.

A bheur hag's skin is blue-white, like that of a person who has frozen to death. Her hair is pale white, and she is emaciated, as if she were a person who had survived winter by eating bark and leather. Her eyes are pale and surrounded by dark, bruise-colored flesh. A bheur carries a twisted gray wooden staff, which she can ride like a flying broom and augments her magical powers.

## Cold Hearts

Bheur hags are attracted to selfish actions justified by deadly cold, such as murdering a traveler for a winter coat, chopping down a dryad's grove for firewood, and so on. These actions are especially sweet to a bheur if they are unwarranted, such as a greedy merchant hoarding more food for the winter than he could possibly eat while others starve. Bheurs love to seed such ideas and thoughts in mortals. They use their ability to manipulate weather to batter villages with snow and freezing cold, hoping to instill despair that turns the villagers against each other.

A bheur hag loves watching unprepared people suffer and die for their mistakes during the winter. She is delighted when mortals make petty, pathetic attempts to survive, such as eating boots and leather scraps when no real food is to be found.

## Awful to Behold

When a bheur hag is fully in the throes of combat and has recently slain one of her foes, she often forgoes a direct attack on her remaining enemies and instead takes a moment to feed on the corpse, dismembering it and tearing meat from bone. The sight of this savagery is enough to render witnesses temporarily insane.

## Covens

A bheur hag that is part of a coven (see the "Hag Covens" sidebar in the Monster Manual) has a challenge rating of 9 (5,000 XP).

```statblock
"name": "Bheur Hag (VGM)"
"size": "Medium"
"type": "fey"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "91"
"hit_dice": "14d8 + 28"
"stats":
- !!int "13"
- !!int "16"
- !!int "14"
- !!int "12"
- !!int "13"
- !!int "16"
"speed": "30 ft."
"saves":
  "Wisdom": !!int "4"
"skillsaves":
  "Nature": !!int "4"
  "Stealth": !!int "6"
  "Perception": !!int "4"
  "Survival": !!int "4"
"damage_immunities": "cold"
"senses": "darkvision 60 ft., passive Perception 14"
"languages": "Auran, Common, Giant"
"cr": "7"
"traits":
- "desc": "The hag's innate spellcasting ability is Charisma (spell save DC 14, dice:\
    \ d20+6 (+6) to hit with spell attacks). She can innately cast the following\
    \ spells, requiring no material components:\n\nAt will: [hold person](/3-Mechanics/CLI/spells/hold-person.md),\
    \ [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md)\n\n1/day each: [control\
    \ weather](/3-Mechanics/CLI/spells/control-weather.md)\n\n3/day each: [cone\
    \ of cold](/3-Mechanics/CLI/spells/cone-of-cold.md), [ice storm](/3-Mechanics/CLI/spells/ice-storm.md),\
    \ [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)"
  "name": "Innate Spellcasting"
- "desc": "The hag carries a graystaff, a length of gray wood that is a focus for\
    \ her inner power. She can ride the staff as if it were a [broom of flying](/3-Mechanics/CLI/items/broom-of-flying.md).\
    \ While holding the staff, she can cast additional spells with her Innate Spellcasting\
    \ trait (these spells are marked with an asterisk). If the staff is lost or destroyed,\
    \ the hag must craft another, which takes a year and a day. Only a bheur hag can\
    \ use a graystaff."
  "name": "Graystaff Magic"
- "desc": "The hag can move across and climb icy surfaces without needing to make\
    \ an ability check. Additionally, difficult terrain composed of ice or snow doesn't\
    \ cost her extra moment."
  "name": "Ice Walk"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d8 + 1|text(10) (2d8 + 1) bludgeoning damage plus dice:1d6|text(3)\
    \ (1d6) cold damage."
  "name": "Slam"
- "desc": "The hag feasts on the corpse of one enemy within 5 feet of her that died\
    \ within the past minute. Each creature of the hag's choice that is within 60\
    \ feet of her and able to see her must succeed on a DC 15 Wisdom saving throw\
    \ or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) of her for\
    \ 1 minute. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, a creature is [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated),\
    \ can't understand what others say, can't read, and speaks only in gibberish;\
    \ the DM controls the creature's movement, which is erratic. A creature can repeat\
    \ the saving throw at the end of each of its turns, ending the effect on itself\
    \ on a success. If a creature's saving throw is successful or the effect ends\
    \ for it, the creature is immune to the hag's Maddening Feast for the next 24\
    \ hours."
  "name": "Maddening Feast"
"lair_actions":
- "desc": "The following lair actions are options for grandmothers and powerful aunties.\
    \ Grandmothers usually have three to five lair actions, aunties usually only one\
    \ (if they have any at all). Unless otherwise noted, any lair action that requires\
    \ a creature to make a saving throw uses the save DC of the hag's most powerful\
    \ ability."
  "name": ""
- "desc": "On initiative count 20 (losing initiative ties), the hag can take a lair\
    \ action to cause one of the following effects, but can't use the same effect\
    \ two rounds in a row:"
  "name": ""
- "desc": "- Until initiative count 20 on the next round, the hag can pass through\
    \ solid walls, doors, ceilings, and floors as if the surfaces weren't there. \
    \ \n- The hag targets any number of doors and windows that she can see, causing\
    \ each one to either open or close as she wishes. Closed doors can be magically\
    \ locked (requiring a successful DC 20 Strength check to force open) until she\
    \ chooses to make them unlocked, or until she uses this lair action again to open\
    \ them.  "
  "name": ""
- "desc": "A powerful bheur hag might have the following additional lair action:"
  "name": ""
- "desc": "- The hag creates a blizzard in a 40-foot-high, 20-foot radius cylinder\
    \ centered on a point she can see within 120 feet of her. The effect lasts until\
    \ initiative count 20 on the next round. The blizzard lightly obscures every creature\
    \ and object in the area for the duration. A creature that enters the blizzard\
    \ for the first time on a turn or starts its turn there is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded)\
    \ until initiative count 20 on the next round.  "
  "name": ""
"regional_effects":
- "desc": "Each hag's lair is the source of three to five regional effects; the home\
    \ of a grandmother, an auntie, or a coven has more effects than the lair of a\
    \ single hag, including some that can directly harm intruders. Any regional effect\
    \ that requires a creature to make a saving throw uses the save DC of the hag's\
    \ most powerful ability. These effects either end immediately if the hag dies\
    \ or abandons the lair, or take up to dice: 2d10|avg|noform (2d10) days to\
    \ fade away."
  "name": ""
- "desc": "The region within 1 mile of a grandmother hag's lair is warped by the creature's\
    \ fell magic, which creates one or more of the following effects:"
  "name": ""
- "desc": "- Birds, rodents, snakes, spiders, or toads (or some other creatures appropriate\
    \ to the hag) are found in great profusion.  \n- Beasts that have an Intelligence\
    \ score of 2 or lower are [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ by the hag and directed to be aggressive toward intruders in the area.  \n-\
    \ Strange carved figurines, twig fetishes, or rag dolls magically appear in trees.\
    \  "
  "name": ""
- "desc": "A powerful bheur hag creates one or more of the following additional regional\
    \ effects within 1 mile of her lair:"
  "name": ""
- "desc": "- Small avalanches of snow intermittently fall, blocking a path or burying\
    \ intruders. A buried creature is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ and has to hold its breath until it is dug out.  \n- Human-sized blocks of ice\
    \ appear, containing frozen corpses. These corpses might break free and attack\
    \ as zombies, or their spirits might attack as specters.  \n- Blizzards come without\
    \ warning. A blizzard occurs once every dice: 2d12|avg|noform (2d12) hours\
    \ and lasts dice: 1d3|avg|noform (1d3) hours. During a storm, creatures moving\
    \ overland travel at half normal speed, and normal visibility is reduced to 30\
    \ feet.  \n- Roads, paths, and trails twist and turn back on themselves, making\
    \ navigation in the area exceedingly difficult.  "
  "name": ""
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Bheur%20Hag.webp"
```
^statblock

```encounter-table
name: Bheur Hag
creatures:
 - 1: Bheur Hag
```

## Environment

arctic