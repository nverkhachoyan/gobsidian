---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/13
- monster/environment/underdark
- monster/size/large
- monster/type/fiend
aliases: ["Devourer"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 138
---
# [Devourer](3-Mechanics\CLI\bestiary\fiend/devourer-vgm.md)
*Source: Volo's Guide to Monsters p. 138*  

Of all the abominations Orcus has unleashed, devourers are among the most feared. These tall, mummy-like fiends wander the planes, consuming souls and, by example, spreading Orcus's creed of replacing all life with everlasting death.

## Instruments of Orcus

A lesser demon that proves itself to Orcus might be granted the privilege of becoming a devourer. The Prince of Undeath transforms such a demon into an 8-foot-tall, desiccated humanoid with a hollowed-out ribcage, then fills the new creature with a hunger for souls. Orcus grants each new devourer the essence of a less fortunate demon to power the devourer's first foray into the planes. Most devourers remain in the Abyss, or on the Astral or Ethereal Plane, pursuing Orcus's schemes and interests in those realms. When Orcus sends devourers to the Material Plane, he often sets them on a mission to create, control, and lead a plague of undead. Skeletons, zombies, ghouls and ghasts, and shadows are particularly attracted to the presence of a devourer.

I've heard of "rib-sticking," but this is ridiculous.

-Volo

## Tormentors of Souls

Devourers hunt humanoids, with the intent of consuming them body and soul. After a devourer brings a target to the brink of death, it pulls the victim's body in and traps the creature within its own ribcage. As the victim tries to stave off death (usually without success), the devourer tortures its soul with telepathic noise. When the victim expires, it undergoes a horrible transformation, springing forth from the devourer's body to begin its new existence as an undead servitor of the monster that spawned it.

## Fiendish Nature

A devourer doesn't require air, food (other than souls), drink, or sleep.

```statblock
"name": "Devourer (VGM)"
"size": "Large"
"type": "fiend"
"alignment": "Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "178"
"hit_dice": "17d10 + 85"
"stats":
- !!int "20"
- !!int "12"
- !!int "20"
- !!int "13"
- !!int "10"
- !!int "16"
"speed": "30 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "Abyssal, telepathy 120 ft."
"cr": "13"
"actions":
- "desc": "The devourer makes two claw attacks and can use either Imprison Soul or\
    \ Soul Rend."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 5|text(12) (2d6 + 5) slashing damage plus dice:6d6|text(21)\
    \ (6d6) necrotic damage."
  "name": "Claw"
- "desc": "The devourer chooses a living humanoid with 0 hit points that it can see\
    \ within 30 feet of it. That creature is teleported inside the devourer's ribcage\
    \ and imprisoned there. A creature imprisoned in this manner has disadvantage\
    \ on death saving throws. If it dies while imprisoned, the devourer regains 25\
    \ hit points, immediately recharges Soul Rend, and gains an additional action\
    \ on its next turn. Additionally, at the start of its next turn, the devourer\
    \ regurgitates the slain creature as a bonus action, and the creature becomes\
    \ an undead. If the victim had 2 or fewer Hit Dice, it becomes a zombie. if it\
    \ had 3 to 5 Hit Dice, it becomes a ghoul. Otherwise, it becomes a wight. A devourer\
    \ can imprison only one creature at a time."
  "name": "Imprison Soul"
- "desc": "The devourer creates a vortex of life-draining energy in a 20-foot radius\
    \ centered on itself. Each humanoid in that area must make a DC 18 Constitution\
    \ saving throw, taking dice:8d10|text(44) (8d10) necrotic damage on a failed\
    \ save, or half as much damage on a successful one. Increase the damage by 10\
    \ for each living humanoid with 0 hit points in that area."
  "name": "Soul Rend (Recharge 6)"
"source":
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Devourer.webp"
```
^statblock

```encounter-table
name: Devourer
creatures:
 - 1: Devourer
```

## Environment

underdark