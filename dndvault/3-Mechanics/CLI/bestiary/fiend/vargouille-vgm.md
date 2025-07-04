---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/1
- monster/environment/desert
- monster/environment/swamp
- monster/environment/underdark
- monster/size/tiny
- monster/type/fiend
aliases: ["Vargouille"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 195
---
# [Vargouille](3-Mechanics\CLI\bestiary\fiend/vargouille-vgm.md)
*Source: Volo's Guide to Monsters p. 195*  

Shrieking, flapping, and hideous to behold-with a body like a severed head and bat-like wings in place of ears—vargouilles boil out of the Abyss to infest other planes of existence, such as Carceri, where they are a menace. Each vargouille carries a disease that creates more of its kind; a flock of vargouilles on the wing is a plague of chaos and evil waiting to happen.

## Abyssal Nuisances

Swarms of vargouilles flap through the caverns and skies of the Abyss. They are given little regard by powerful and intelligent demons since vargouilles can do them no harm. Even the weakest demon, such as a manes or a dretch, fears vargouilles only if they appear in great numbers. In the Lower Planes, vargouilles rarely get the chance to eat live prey other than vermin. More often, they lap up the ichor left behind when one fiend kills another.

The kiss of a vargouille infects a humanoid with a fiendish curse. If allowed to run its course, the curse brings about a gruesome transformation as an abyssal spirit invades the person's body. Over a period of hours, the victim's head takes on fiendish aspects such as fangs, tentacles, and horns. At the same time, the person's ears grow larger, expanding and transforming into wing-like appendages. In the final moments, the victim's head tears away from the body in a fountain of blood, becoming another vargouille, which often then eagerly laps up its own life fluids. Sunlight or the brilliant illumination of a [daylight](/3-Mechanics/CLI/spells/daylight.md) spell can delay this transformation, and vargouilles instinctively shun bright light as a result.

## The World Awaits

Because of their instinctive hunger for living prey, vargouilles are eager to escape the Lower Planes. On rare occasions, the summoning of a demon to another plane can bring a vargouille along for the ride, attaching itself like a tick. The precautions a mortal takes to contain and control a summoned demon rarely account for a stowaway, and thus a vargouille enters the world unbidden.

## Ghastly Reproduction

Vargouilles that roam free on the Material Plane are a dire threat to all creatures, especially humanoids. Their awful shrieking can paralyze other creatures with fear, and such victims are helpless to resist a vargouille's accursed kiss.

```statblock
"name": "Vargouille (VGM)"
"size": "Tiny"
"type": "fiend"
"alignment": "Chaotic Evil"
"ac": !!int "12"
"hp": !!int "13"
"hit_dice": "3d4 + 6"
"stats":
- !!int "6"
- !!int "14"
- !!int "14"
- !!int "4"
- !!int "7"
- !!int "2"
"speed": "5 ft., fly 40 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "understands Abyssal, Infernal, and any languages it knew before becoming\
  \ a vargouille but can't speak"
"cr": "1"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage plus dice:3d6|text(10)\
    \ (3d6) poison damage."
  "name": "Bite"
- "desc": "The vargouille kisses one [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ humanoid within 5 feet of it. The target must succeed on a DC 12 Charisma saving\
    \ throw or become cursed. The cursed target loses 1 point of Charisma after each\
    \ hour, as its head takes on fiendish aspects. The curse doesn't advance while\
    \ the target is in sunlight or the area of a [daylight](/3-Mechanics/CLI/spells/daylight.md)\
    \ spell; don't count that time. When the cursed target's Charisma becomes 2, it\
    \ dies, and its head tears from its body and becomes a new vargouille. Casting\
    \ [remove curse](/3-Mechanics/CLI/spells/remove-curse.md), [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md),\
    \ or a similar spell on the target before the transformation is complete can end\
    \ the curse. Doing so undoes the changes made to the target by the curse."
  "name": "Kiss"
- "desc": "The vargouille shrieks. Each humanoid and beast within 30 feet of the vargouille\
    \ and able to hear it must succeed on a DC 12 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of the vargouille's next turn. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, a target is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned).\
    \ If a target's saving throw is successful or the effect ends for it, the target\
    \ is immune to the Stunning Shriek of all vargouilles for 1 hour."
  "name": "Stunning Shriek"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Vargouille.webp"
```
^statblock

```encounter-table
name: Vargouille
creatures:
 - 1: Vargouille
```

## Environment

underdark, swamp, desert