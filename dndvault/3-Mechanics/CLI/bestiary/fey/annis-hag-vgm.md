---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/6
- monster/environment/hill
- monster/environment/mountain
- monster/size/large
- monster/type/fey
aliases: ["Annis Hag"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 159
---
# [Annis Hag](3-Mechanics\CLI\bestiary\fey/annis-hag-vgm.md)
*Source: Volo's Guide to Monsters p. 159*  

Annis hags lair in mountains or hills. Despite being hunchbacked and hump-shouldered, they are the largest and most physically imposing of their kind, standing eight feet tall.

## Tormenting the Weak

Although annis hags can easily tear a grown man apart, they love hunting children, preferring their flesh above all others. They use the flayed skin of such victims to make supple leather, and a hag's lair often shows the signs of this industry.

Annis hags leave tokens of their cruelty at the edges of forests and other areas they claim. In this way, they provoke fear and paranoia in nearby villages and settlements. To an annis hag, nothing is sweeter than turning a vibrant community into a place paralyzed with terror, where folk never venture out at night, strangers are met with suspicion and anger, and parents warn their children to "be good, or the annis will get you."

## Child Corrupter

When an annis feels especially cruel, she disguises herself as a kindly-looking elderly woman, approaches a child in a remote place, and gives it an iron token that it can use to confide in her. Over time, "Granny" convinces the child that it's okay to have bad thoughts and do bad deeds-starting with breaking things or wandering outside without permission, then graduating to pushing someone down the stairs or setting a house on fire. Sooner or later, the child's family and community become terrified of the "bad seed" and must face the awful decision of whether the child should be punished or exiled.

## Tribe Mother

Much in the way that they befriend children in order to corrupt them, annis hags have a tendency for adopting a group of ogres, trolls, or other loutish creatures, ruling them through brute strength, verbal abuse, and superstition.

## Covens

An annis hag that is part of a coven (see the "Hag Covens" sidebar in the Monster Manual) has a challenge rating of 8 (3,900 XP).

## Iron Token

An annis hag can pull out one of her iron teeth or nails and spend 1 minute shaping and polishing it into the form of a coin, a ring, or a tiny mirror. Thereafter, any creature that holds this iron token can have a whispered conversation with the hag, provided the creature and the hag are on the same plane of existence and within 10 miles of each other. The holder of the token can hear only the hag's voice, not those of any other creatures or any ambient noise around the hag. Similarly, the hag can hear the holder of the token and not the noise around it.

A hag can have up to three iron tokens active at one time. As an action, she can discern the direction and approximate distance to all of her active tokens. She can instantaneously deactivate any of her tokens at any distance (no action required), whereupon the token retains its current form but loses its magical properties.

```statblock
"name": "Annis Hag (VGM)"
"size": "Large"
"type": "fey"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "75"
"hit_dice": "10d10 + 20"
"stats":
- !!int "21"
- !!int "12"
- !!int "14"
- !!int "13"
- !!int "14"
- !!int "15"
"speed": "40 ft."
"saves":
  "Constitution": !!int "5"
"skillsaves":
  "Deception": !!int "5"
  "Perception": !!int "5"
"damage_resistances": "cold; bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "Common, Giant, Sylvan"
"cr": "6"
"traits":
- "desc": "The hag's innate spellcasting ability is Charisma (spell save DC 13). She\
    \ can innately cast the following spells:\n\n3/day each: [disguise self](/3-Mechanics/CLI/spells/disguise-self.md)\
    \ (including the form of a Medium humanoid), [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md)"
  "name": "Innate Spellcasting"
"actions":
- "desc": "The annis makes three attacks: one with her bite and two with her claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 5|text(15) (3d6 + 5) slashing damage."
  "name": "Claw"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:9d6 + 5|text(36) (9d6 + 5) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 15)\
    \ if it is a Large or smaller creature. Until the grapple ends, the target takes\
    \ dice:9d6 + 5|text(36) (9d6 + 5) bludgeoning damage at the start of each\
    \ of the hag's turns. The hag can't make attacks while grappling a creature in\
    \ this way."
  "name": "Crushing Hug"
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
- "desc": "A powerful annis hag might have the following additional lair action:"
  "name": ""
- "desc": "- The hag creates a thick cloud of caustic black smoke that fills a 20-foot-radius\
    \ sphere centered on a point she can see within 120 feet her. The cloud lasts\
    \ until initiative count 20 on the next round. Creatures and objects in or behind\
    \ the smoke are heavily obscured. A creature that enters the cloud for the first\
    \ time on a turn or starts its turn there takes dice:3d6|text(10) (3d6) acid\
    \ damage.  "
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
- "desc": "A powerful annis hag creates one or more of the following additional regional\
    \ effects within 1 mile of her lair:"
  "name": ""
- "desc": "- The gravel stones on a safe-looking path, road, or trails occasionally\
    \ become sharp for 100-foot intervals. Walking on these areas is like walking\
    \ on caltrops.  \n- Small avalanches of rock intermittently fall, blocking a path\
    \ or burying intruders. A buried creature is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ and has to hold its breath until it is dug out.  \n- Strange laughter, sounding\
    \ like that of children or the hag herself, occasionally pierces the silence.\
    \  \n- Small cairns appear along the route of travelers, containing anything from\
    \ mysterious bones to nothing at all. These cairns might be haunted by skeletons,\
    \ specters, or hostile fey.  "
  "name": ""
"source":
- "VGM"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Annis%20Hag.webp"
```
^statblock

```encounter-table
name: Annis Hag
creatures:
 - 1: Annis Hag
```

## Environment

mountain, hill