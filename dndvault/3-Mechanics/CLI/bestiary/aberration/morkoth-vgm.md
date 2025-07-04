---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/11
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/aberration
aliases: ["Morkoth"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 177
---
# [Morkoth](3-Mechanics\CLI\bestiary\aberration/morkoth-vgm.md)
*Source: Volo's Guide to Monsters p. 177*  

Ancient and devious, morkoths are voracious collectors. Each one travels the planes, amassing the valuables, oddities, and castoffs of the multiverse to make its collection ever more complete.

## Spawned by a God

Long ago, a deity of greed and strife perished in the battles among the immortals. Its body drifted through the Astral Plane, eventually becoming a petrified husk. This corpse floated up against a pearlescent remnant of celestial matter imbued with life and life-giving magic. The collision shattered both objects and released a storm of chaotic energy. Countless islands of mixed matter spun away into the silvery void. Within some of them, a vein of pearl-like material held a bit of the deity's rejuvenated supernatural vitality, which spontaneously created a habitable environment. On those same islands, bits of the god's petrified flesh came back to life, in the form of tentacled monstrosities brimming with malice and greed. Ever since that time, each morkoth has had an extraplanar island to call home.

The pearly matter inside an island enables it to glide on planar currents, maintains the island's environment, and keeps the place safe from harmful external effects. A morkoth's island might be found anywhere from the bottom of the ocean to the void of the Astral Plane. One could float in the skies of Avernus in the Nine Hells without being destroyed and without causing harm to its residents. Whatever is on or within a certain distance of a morkoth's isle travels with it in its journey through the planes. Thus, people from lost civilizations and creatures or objects from bygone ages might be found within a morkoth's dominion.

Some islands travel a specific route, arriving at the same destinations regularly over a cycle of years. Others are tied to a particular place or group of locales, and still others move erratically through the cosmos. Rarely, a morkoth learns to control its island's movement, so the island goes wherever its master wishes.

A morkoth spends its time watching over its collection and plotting to acquire more possessions. The monster hoards vast stores of treasure and knowledge. Its island holds numerous captives, which it considers part of its collection. Some inhabitants, such as descendants of original prisoners, might view the morkoth as a ruler or a god. A morkoth's storehouse of wealth and lore attracts would-be plunderers, of course, as well as those seeking something specific the morkoth has or knows. The creature shows no mercy to those that try to steal from it, but it can be bargained with by a visitor that offers the morkoth something it desires.

No morkoth freely gives away what it owns. Morkoths exist to acquire, and they give up possessions only if doing so helps their hoard grow.

A morkoth knows every object in its collection and can track its possessions through the planes. Someone who dares to steal from a morkoth, or breaks a deal with one, will know no rest until the morkoth is slain or all promises are kept.

Collectors of everything odd, unusual, and valuable-hopefully not including you.

-Volo

## No Rhyme or Reason

A morkoth's island has the qualities of a dreamscape in which nature and predictability take a back seat to strangeness and chaos. Upon it is a jumble of objects and a mixture of creatures, some of which date from forgotten times. An island might have natural-looking illumination, but most are shrouded in twilight, and on any of them, mists and shadows can appear without notice. The environment is warm and wet, a subtropical or tropical climate that keeps the morkoth and its "guests" comfortable.

## Primeval Hoarders

Morkoths are driven by greed and selfishness, mixed with a yearning for conflict. They desire anything they don't possess, have no scruples about taking what they crave, and endeavor to keep everything they collect.

## A Morkoth's Lair

A morkoth claims dominion over an entire island, and it also maintains a central sanctum on that isle. This lair is most often a twisted network of narrow tunnels that connect several underground chambers, although other structural forms might be incorporated. The morkoth dwells among its most prized possessions in a spacious vault at the core of the warren, where the pearly matter of the island is also located. Sections of the lair and its center might be kept dry to better protect and preserve collected objects and creatures, but most of the lair is underwater.

A morkoth encountered in its lair has a challenge rating of 12 (8,400 XP).

```statblock
"name": "Morkoth (VGM)"
"size": "Medium"
"type": "aberration"
"alignment": "Chaotic Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "130"
"hit_dice": "20d8 + 40"
"stats":
- !!int "14"
- !!int "14"
- !!int "14"
- !!int "20"
- !!int "15"
- !!int "13"
"speed": "25 ft., swim 50 ft."
"saves":
  "Dexterity": !!int "6"
  "Wisdom": !!int "6"
  "Intelligence": !!int "9"
"skillsaves":
  "Stealth": !!int "6"
  "Perception": !!int "10"
  "History": !!int "9"
  "Arcana": !!int "9"
"damage_resistances": "bludgeoning, piercing, slashing from nonmagical attacks"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 20"
"languages": "telepathy 120 ft."
"cr": "11"
"traits":
- "desc": "The morkoth is an 11th-level spellcaster. Its spellcasting ability is Intelligence\
    \ (save DC 17, dice: d20+9 (+9) to hit with spell attacks). The morkoth has\
    \ the following wizard spells prepared:\n\nCantrips (at will): [acid splash](/3-Mechanics/CLI/spells/acid-splash.md),\
    \ [mage hand](/3-Mechanics/CLI/spells/mage-hand.md), [mending](/3-Mechanics/CLI/spells/mending.md),\
    \ [ray of frost](/3-Mechanics/CLI/spells/ray-of-frost.md), [shocking grasp](/3-Mechanics/CLI/spells/shocking-grasp.md)\n\
    \n1st level (4 slots): [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [identify](/3-Mechanics/CLI/spells/identify.md), [shield](/3-Mechanics/CLI/spells/shield.md),\
    \ [witch bolt](/3-Mechanics/CLI/spells/witch-bolt.md)\n\n2nd level (3 slots):\
    \ [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md),\
    \ [shatter](/3-Mechanics/CLI/spells/shatter.md)\n\n3rd level (3 slots): [dispel\
    \ magic](/3-Mechanics/CLI/spells/dispel-magic.md), [lightning bolt](/3-Mechanics/CLI/spells/lightning-bolt.md),\
    \ [sending](/3-Mechanics/CLI/spells/sending.md)\n\n4th level (3 slots): [dimension\
    \ door](/3-Mechanics/CLI/spells/dimension-door.md), [Evard's black tentacles](/3-Mechanics/CLI/spells/evards-black-tentacles.md)\n\
    \n5th level (3 slots): [geas](/3-Mechanics/CLI/spells/geas.md), [scrying](/3-Mechanics/CLI/spells/scrying.md)\n\
    \n6th level (1 slots): [chain lightning](/3-Mechanics/CLI/spells/chain-lightning.md)"
  "name": "Spellcasting"
- "desc": "The morkoth can breathe air and water."
  "name": "Amphibious"
"actions":
- "desc": "The morkoth makes three attacks: two with its bite and one with its tentacles\
    \ or three with its bite."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 2|text(9) (2d6 + 2) slashing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 15 ft., one target.\
    \ Hit: dice:3d8 + 2|text(15) (3d8 + 2) bludgeoning damage, and the target\
    \ is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape DC 14)\
    \ if it is a Large or smaller creature. Until this grapple ends. the target is\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) and takes dice:3d8\
    \ + 2|text(15) (3d8 + 2) bludgeoning damage at the start of each of the morkoth's\
    \ turns. and the morkoth can't use its tentacles on another target."
  "name": "Tentacles"
- "desc": "The morkoth projects a 30-foot cone of magical energy. Each creature in\
    \ that area must make a DC 17 Wisdom saving throw. On a failed save, the creature\
    \ is [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) by the morkoth for\
    \ 1 minute. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) in this\
    \ way, the target tries to get as close to the morkoth as possible, using its\
    \ actions to Dash until it is within 5 feet of the morkoth. A [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ target can repeat the saving throw at the end of each of its turns and whenever\
    \ it takes damage, ending the effect on itself on a success. If a creature's saving\
    \ throw is successful or the effect ends for it, the creature has advantage on\
    \ saving throws against the morkoth's Hypnosis for 24 hours."
  "name": "Hypnosis"
"reactions":
- "desc": "If the morkoth makes a successful saving throw against a spell, or a spell\
    \ attack misses it, the morkoth can choose another creature (including the spellcaster)\
    \ it can see within 120 feet of it. The spell targets the chosen creature instead\
    \ of the morkoth. If the spell forced a saving throw, the chosen creature makes\
    \ its own save. If the spell was an attack, the attack roll is rerolled against\
    \ the chosen creature."
  "name": "Spell Reflection"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), the morkoth takes a lair\
    \ action to cause one of the effects described below:"
  "name": ""
- "desc": "- The morkoth uses its Hypnosis action, originating at a point within 120\
    \ feet of itself. It doesn't need to see the effect's point of origin.  \n- The\
    \ morkoth casts [darkness](/3-Mechanics/CLI/spells/darkness.md), [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md),\
    \ or [misty step](/3-Mechanics/CLI/spells/misty-step.md), using Intelligence as\
    \ its spellcasting ability and without expending a spell slot.  "
  "name": ""
"regional_effects":
- "desc": "The island surrounding a morkoth's lair is warped by the creature's presence,\
    \ creating the following effects:"
  "name": ""
- "desc": "- The morkoth is aware of any new arrival, whether an object or a creature,\
    \ on its island or in its sanctum. As an action, the morkoth can locate any one\
    \ creature or object on the island. Visitors to the island feel as though they\
    \ are being watched, even when they aren't.  \n- Each time a creature that has\
    \ been on the island for less than a year finishes a short or long rest, it must\
    \ make a DC 10 Intelligence ([Investigation](/3-Mechanics/CLI/rules/skills.md#Investigation))\
    \ check. On a failure, the creature has misplaced one possession (chosen by the\
    \ player, if the creature is that player's character). The possession remains\
    \ nearby but concealed for a short time, so it can be recovered with a successful\
    \ DC 15 Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) check.\
    \ An object that is misplaced but not recovered ends up in the morkoth's lair\
    \ 1 hour later. If the creature later goes to the morkoth's lair, its lost possessions\
    \ stand out in its perception and are easily recovered.  \n- Entrances to the\
    \ morkoth's lair have an enchantment that the morkoth can activate or suppress\
    \ at any time while it's in its lair and not [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated).\
    \ Any creature within 30 feet of such an entrance and able to see it must make\
    \ a DC 15 Wisdom saving throw. On a failed save, the creature feels an intense\
    \ urge to use its movement on each of its turns to enter the lair and to move\
    \ toward the morkoth's location (the target doesn't realize it's heading toward\
    \ a creature). The target moves toward the morkoth by the most direct route. As\
    \ soon as it can see the morkoth, the target can repeat the saving throw, ending\
    \ the effect on itself on a success. It can also repeat the saving throw at the\
    \ end of each of its turns and every time it takes damage.  \n- With a thought\
    \ (no action required), the morkoth can initiate a change in the water within\
    \ its lair that takes effect 1 minute later. The water can be as breathable and\
    \ clear as air, or it can be normal water (ranging in clarity from murky to clear).\
    \  "
  "name": ""
- "desc": "If the morkoth dies, these regional effects end immediately."
  "name": ""
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Morkoth.webp"
```
^statblock

```encounter-table
name: Morkoth
creatures:
 - 1: Morkoth
```

## Environment

underwater, coastal