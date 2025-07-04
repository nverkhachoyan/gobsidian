---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/11
- monster/environment/coastal
- monster/environment/underwater
- monster/size/large
- monster/type/elemental
aliases: ["Marid"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Monster Manual p. 146, Dragon of Icespire Peak, Storm Lord's Wrath, Tasha's Cauldron of Everything
---
# [Marid](3-Mechanics\CLI\bestiary\elemental/marid.md)
*Source: Monster Manual p. 146, Dragon of Icespire Peak, Storm Lord's Wrath, Tasha's Cauldron of Everything*  

Hailing from the Elemental Plane of Water, the marids are the most wondrous of genie-kind. Although all genies wield great power, even the lowliest marid sees itself as clearly superior to the flighty djinn, the ground-hugging dao, and the fuming efreet. Large and piscine, marids are a strange sight to behold, particularly when clad in the finely stitched vests and colorful pantaloons they favor. They speak in voices as soft as the sea breeze or as sonorous as storm waves breaking against a rocky cliff. In flight, their lower bodies transform into columns of foamy water.

## Water Lords

Water is a marid's native element, and the genie can manipulate water in virtually any way it desires. A marid can walk on water and breathe naturally beneath its surface. It can create water or shape clouds of fog and mist from the vapor in the air. It can even transform itself into mist, or use water as a weapon to bludgeon its foes.

## Marid Homes

Marids are rare on the Material Plane. They inhabit mighty and majestic coral fortresses located in the Elemental Plane of Water. These citadels float in the depths of the plane and contain opulent, air-filled chambers where slaves and guests reside.

A marid doesn't expect much from its slaves, simply wanting to have them for the status of ownership. Marids go out of their way to obtain skilled slaves, and aren't above kidnapping mortal artists, entertainers, or storytellers for use in their courts.

## Egotistical Hierarchs

All marids claim a title of nobility, and the race is awash in shahs, sultans, muftis, and khedives. Most of these titles are mere pretense on the part of the self-important marids.

Marids treat all others-including other genies-as inferiors of various grades, ranging from poor cousins to petty annoyances. They tolerate djinn, dislike dao, and despise efreet.

Humanoids are among the lowest of the creatures that marids must tolerate, although they sometimes deal with powerful wizards and exceptional leaders on an almost-equal footing. Doing so has sometimes proven to be a mistake, since wizards have managed to imprison marids in conch shells, flasks, and decanters over the ages. Bribery and flattery are the best means of dealing with marids, to which an obsequious mortal is a creature that knows its place.

## Whimsical Storytellers

Marids are champion tale-tellers, whose favorite legends emphasize the prowess of marids in general and of the speaker in particular. Fanciful genies, they lie often and creatively. They aren't always malicious in their deception, but embellishments suit their fancy. Marids consider it a crime for a lesser being to interrupt one of their tales, and offending a marid is a sure way to invoke its wrath.

> [!quote]- A quote from Kesto Brighteyes, Gnome Proprietor of the Parted Veil, a bookshop in Sigil  
> 
> The marid poured out of the flask like water and said, 'Your wish is my command.' The halfling, overjoyed, wished for immortality, so the marid polymorphed him into a fish that flopped around humorously until, finally, it expired. It's a cautionary tale that has survived through the ages, so I suppose the halfling got his wish.


```statblock
"name": "Marid"
"size": "Large"
"type": "elemental"
"alignment": "Chaotic Neutral"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "229"
"hit_dice": "17d10 + 136"
"stats":
- !!int "22"
- !!int "12"
- !!int "26"
- !!int "18"
- !!int "17"
- !!int "18"
"speed": "30 ft., fly 60 ft., swim 90 ft."
"saves":
  "Charisma": !!int "8"
  "Dexterity": !!int "5"
  "Wisdom": !!int "7"
"damage_resistances": "acid, cold, lightning"
"senses": "blindsight 30 ft., darkvision 120 ft., passive Perception 13"
"languages": "Aquan"
"cr": "11"
"traits":
- "desc": "The marid's innate spellcasting ability is Charisma (spell save DC 16,\
    \ dice: d20+8 (+8) to hit with spell attacks). It can innately cast the following\
    \ spells, requiring no material components:\n\nAt will: [create or destroy\
    \ water](/3-Mechanics/CLI/spells/create-or-destroy-water.md), [detect evil and\
    \ good](/3-Mechanics/CLI/spells/detect-evil-and-good.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [fog cloud](/3-Mechanics/CLI/spells/fog-cloud.md), [purify food and drink](/3-Mechanics/CLI/spells/purify-food-and-drink.md)\n\
    \n1/day each: [conjure elemental](/3-Mechanics/CLI/spells/conjure-elemental.md)\
    \ ([water elemental](/3-Mechanics/CLI/bestiary/elemental/water-elemental.md) only),\
    \ [control water](/3-Mechanics/CLI/spells/control-water.md), [gaseous form](/3-Mechanics/CLI/spells/gaseous-form.md),\
    \ [invisibility](/3-Mechanics/CLI/spells/invisibility.md), [plane shift](/3-Mechanics/CLI/spells/plane-shift.md)\n\
    \n3/day each: [tongues](/3-Mechanics/CLI/spells/tongues.md), [water breathing](/3-Mechanics/CLI/spells/water-breathing.md),\
    \ [water walk](/3-Mechanics/CLI/spells/water-walk.md)"
  "name": "Innate Spellcasting"
- "desc": "The marid can breathe air and water."
  "name": "Amphibious"
- "desc": "If the marid dies, its body disintegrates into a burst of water and foam,\
    \ leaving behind only equipment the marid was wearing or carrying."
  "name": "Elemental Demise"
"actions":
- "desc": "The marid makes two trident attacks."
  "name": "Multiattack"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+10 (+10) to hit, reach 5\
    \ ft. or range 20/60 ft., one target. Hit: dice:2d6 + 6|text(13) (2d6 + 6)\
    \ piercing damage, or dice:2d8 + 6|text(15) (2d8 + 6) piercing damage if used\
    \ with two hands to make a melee attack."
  "name": "Trident"
- "desc": "The marid magically shoots water in a 60-foot line that is 5 feet wide.\
    \ Each creature in that line must make a DC 16 Dexterity saving throw. On a failure,\
    \ a target takes dice:6d6|text(21) (6d6) bludgeoning damage and, if it is\
    \ Huge or smaller, is pushed up to 20 feet away from the marid and knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \ On a success, a target takes half the bludgeoning damage, but is neither pushed\
    \ nor knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Water Jet"
"source":
- "MM"
- "WDMM"
- "GoS"
- "DIP"
- "SLW"
- "EGW"
- "TCE"
- "JttRC"
- "DSotDQ"
- "SatO"
- "BMT"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Marid.webp"
```
^statblock

```encounter-table
name: Marid
creatures:
 - 1: Marid
```

## Environment

underwater, coastal