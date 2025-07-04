---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/forest
- monster/environment/hill
- monster/environment/swamp
- monster/size/small
- monster/type/fey
aliases: ["Redcap"]
NoteIcon: monster
BestiaryType: fey
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 188
---
# [Redcap](3-Mechanics\CLI\bestiary\fey/redcap-vgm.md)
*Source: Volo's Guide to Monsters p. 188*  

A redcap is a homicidal fey creature born of blood lust. Redcaps, although small, have formidable strength, which they use to hunt and kill without hesitation or regret.

## Blood Lust Personified

In the Feywild, or where that plane touches the world at a fey crossing, if a sentient creature acts on an intense desire for bloodshed, one or more redcaps might appear where the blood of a slain person soaks the ground. At first, new red caps look like tiny bloodstained mushrooms just pushing their caps out of the soil. When moonlight shines on one of these caps, a creature that looks like a wizened and undersized gnome with a hunched back and a sinewy frame springs from the earth. The creature has a pointed leather cap, pants of similar material, heavy iron boots, and a heavy bladed weapon. From the moment it awakens, a redcap desires only murder and carnage, and it sets out to satisfy these cravings.

Redcaps lack subtlety. They live for direct confrontation and the mayhem of mortal combat. Even if a redcap wanted to be stealthy, its iron boots force it to take ponderous, thunderous steps. When a redcap is near to potential prey, though, it can close the distance quickly and get in a vicious swing of its weapon before the target can react.

Also, some redcaps can sense the being whose murderous acts led to their birth. A redcap might use this innate connection to find its creator and make that creature its first victim. Others seek out their maker to enjoy proximity to a kindred spirit. An individual responsible for the creation of multiple redcaps at the same site could attract the entire group to serve as cohorts, emulating that creature's murderous handiwork.

In any case, if a redcap works with another being, the redcap demands to be paid in victims. A patron who tries to stifle a redcap's natural and necessary urge for blood risks becoming the redcap's next target.

As subtle as a flung battleaxe.

-Volo

## Steeped in Slaughter

To sustain its unnatural existence, a redcap has to soak its hat in the fresh blood of its victims. When a redcap is born, its hat is coated with wet blood, and it knows that if the blood isn't replenished at least once every three days, the redcap vanishes as if it had never been. A redcap's desire to slay is rooted in its will to survive.

## Bloodthirsty Mercenaries

Redcaps don't usually operate in groups, but in some circumstances they might be fond in the employ of hags and dark mages that know methods to call redcaps out of the Feywild and put them to work as grisly servants.

```statblock
"name": "Redcap (VGM)"
"size": "Small"
"type": "fey"
"alignment": "Chaotic Evil"
"ac": !!int "13"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "6d6 + 24"
"stats":
- !!int "18"
- !!int "13"
- !!int "18"
- !!int "10"
- !!int "12"
- !!int "9"
"speed": "25 ft."
"skillsaves":
  "Athletics": !!int "6"
  "Perception": !!int "3"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Common, Sylvan"
"cr": "3"
"traits":
- "desc": "While moving, the redcap has disadvantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks."
  "name": "Iron Boots"
- "desc": "While grappling, the redcap is considered to be Medium. Also, wielding\
    \ a heavy weapon doesn't impose disadvantage on its attack rolls."
  "name": "Outsize Strength"
"actions":
- "desc": "The redcap makes three attacks with its wicked sickle."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 4|text(9) (2d4 + 4) slashing damage."
  "name": "Wicked Sickle"
- "desc": "The redcap moves up to its speed to a creature it can see and kicks with\
    \ its iron boots. The target must succeed on a DC 14 Dexterity saving throw or\
    \ take dice:3d10 + 4|text(20) (3d10 + 4) bludgeoning damage and be knocked\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Ironbound Pursuit"
"source":
- "VGM"
- "ToA"
- "BGDIA"
- "IMR"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Redcap.webp"
```
^statblock

```encounter-table
name: Redcap
creatures:
 - 1: Redcap
```

## Environment

forest, swamp, hill