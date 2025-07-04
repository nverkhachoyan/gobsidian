---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/18
- monster/size/tiny
- monster/type/undead
aliases: ["Demilich"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 48
---
# [Demilich](3-Mechanics\CLI\bestiary\undead/demilich.md)
*Source: Monster Manual p. 48*  

The immortality granted to a lich lasts only as long as it feeds mortal souls to its phylactery. If it falters or fails in that task, its bones turn to dust until only its skull remains. This "demilich" contains only a fragment of the lich's malevolent life force-just enough so that if it is disturbed, these remains rise into the air and assume a wraith-like form. The skull then emits a terrifying howl that can slay the weak-hearted and leave others trembling with fear. Left alone, it sinks back down and returns to the empty peace of its existence.

Few liches seek to become demiliches, for it means an end to the existence they hoped to preserve by becoming undead. However, time can erode the lich's reason and memory, causing it to retreat into its ancient tomb and forget to feed on souls. The spells it once knew fade from its mind, and it no longer channels the arcane energy it wielded as a lich. However, even as a mere skull it remains a deadly and vexing enemy.

## Enduring Existence

Even after a lich is reduced to a demilich state, its phylactery survives. As long as its phylactery is intact, the demilich can't be permanently destroyed. Its skull reforms after `dice: 1d10|avg|noform` (`1d10`) days, restoring the creature to its wretched state. If it has the presence of mind to do so, a demilich can reclaim its former power by feeding just one soul to its phylactery. Doing so restores the demilich to lich form, reconstituting its undead body.

## Undead Nature

A demilich doesn't require air, food, drink, or sleep. So great is a demilich's will to survive that it always has the maximum number of hit points for its Hit Dice, instead of average hit points.

## A Demilich's Lair

A demilich hides its earthly remains and treasures in a labyrinthine tomb guarded by monsters and traps. At the heart of this labyrinth rests the demilich's skull and the dust from its other bones. In its crypt, a demilich has access to lair actions and additional uses for its legendary actions. Its whole lair also has unique traits. A demilich in its lair has a challenge rating of 20 (24,500 XP).

## Acererak and His Disciples

The transformation into a demilich isn't a bitter end for all liches that experience it. Made as a conscious choice, the path of the demilich becomes the next step in a dark evolution. The lich Acererak-a powerful wizard and demonologist and the infamous master of the Tomb of Horrors-anticipated his own transformation, preparing for it by setting enchanted gemstones into his skull's eye sockets and teeth. Each of these soul gems possessed the power to capture the souls on which his phylactery would feed.

Acererak abandoned his physical body, accepting that it would molder and dissolve to dust while he traveled the planes as a disembodied consciousness. If the skull that was his last physical remains was ever disturbed, its gems would claim the souls of the insolent intruders to his tomb, magically transferring them to his phylactery.

Liches who follow Acererak's path believe that by becoming free of their bodies, they can continue their quest for power beyond the mortal world. As their patron did, they secure their remains within well-guarded vaults, using soul gems to maintain their phylacteries and destroy the adventurers who disturb their lairs.

Acererak or another demilich like him has a challenge rating of 21 (33,000 XP), or 23 (50,000 XP) in its lair, and gains the following additional action option.

### Trap Soul

The demilich targets one creature that it can see within 30 feet of it. The target must make a DC 19 Charisma saving throw. On a failed save, the target's soul is magically trapped inside one of the demilich's gems. While the soul is trapped, the target's body and all the equipment it is carrying cease to exist. On a successful save, the target takes `dice:7d6|text(24)` (`7d6`) necrotic damage, and if this damage reduces the target to 0 hit points, its soul is trapped as if it failed the saving throw. A soul trapped in a gem for 24 hours is devoured and ceases to exist.

If the demilich drops to 0 hit points, it is destroyed and turns to powder, leaving behind its gems. Crushing a gem releases any soul trapped within, at which point the target's body re-forms in an unoccupied space nearest to the gem and in the same state as when it was trapped.

> [!quote]- A quote from Epitaph of a demilich  
> 
> I, Achererak the Eternal, beckon you to your doom. Come, foolish ones, plunger my treasures, if you dare! Others have tried. All have failed! From your skin, tapestries shall be woven, and your bones will carpet my tomb. Only I am beyond Death's reach. Only I know the secret to true immortality!


```statblock
"name": "Demilich"
"size": "Tiny"
"type": "undead"
"alignment": "Neutral Evil"
"ac": !!int "20"
"ac_class": "natural armor"
"hp": !!int "80"
"hit_dice": "32d4"
"stats":
- !!int "1"
- !!int "20"
- !!int "10"
- !!int "20"
- !!int "17"
- !!int "20"
"speed": "0 ft., fly 30 ft. (hover)"
"saves":
  "Charisma": !!int "11"
  "Wisdom": !!int "9"
  "Intelligence": !!int "11"
  "Constitution": !!int "6"
"damage_resistances": "bludgeoning, piercing, slashing from magic weapons"
"damage_immunities": "necrotic; poison; psychic; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
  \ [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified), [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned),\
  \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "truesight 120 ft., passive Perception 13"
"languages": ""
"cr": "18"
"traits":
- "desc": "If the demilich is subjected to an effect that allows it to make a saving\
    \ throw to take only half damage, it instead takes no damage if it succeeds on\
    \ the saving throw, and only half damage if it fails."
  "name": "Avoidance"
- "desc": "If the demilich fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The demilich is immune to effects that turn undead."
  "name": "Turn Immunity"
"actions":
- "desc": "The demilich emits a bloodcurdling howl. Each creature within 30 feet of\
    \ the demilich that can hear the howl must succeed on a DC 15 Constitution saving\
    \ throw or drop to 0 hit points. On a successful save, the creature is [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ until the end of its next turn."
  "name": "Howl (Recharge 5-6)"
- "desc": "The demilich targets up to three creatures that it can see within 10 feet\
    \ of it. Each target must succeed on a DC 19 Constitution saving throw or take\
    \ dice:6d6|text(21) (6d6) necrotic damage, and the demilich regains hit points\
    \ equal to the total damage dealt to all targets."
  "name": "Life Drain"
"legendary_actions":
- "desc": "The demilich flies up to half its flying speed."
  "name": "Flight"
- "desc": "The demilich magically swirls its dusty remains. Each creature within 10\
    \ feet of the demilich, including around a corner, must succeed on a DC 15 Constitution\
    \ saving throw or be [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) until\
    \ the end of the demilich's next turn. A creature that succeeds on the saving\
    \ throw is immune to this effect until the end of the demilich's next turn."
  "name": "Cloud of Dust"
- "desc": "Each creature with in 30 feet of the demilich must make a DC 15 Constitution\
    \ saving throw. On a failed save, the creature's hit point maximum is magically\
    \ reduced by dice:3d6|text(10) (3d6). If a creature's hit point maximum is\
    \ reduced to 0 by this effect, the creature dies. A creature's hit point maximum\
    \ can be restored with the  [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ spell or similar magic."
  "name": "Energy Drain (Costs 2 Actions)"
- "desc": "The demilich targets one creature it can see within 30 feet of it. The\
    \ target must succeed on a DC 15 Wisdom saving throw or be magically cursed. Until\
    \ the curse ends, the target has disadvantage on attack rolls and saving throws.\
    \ The target can repeat the saving throw at the end of each of its turns, ending\
    \ the curse on a success."
  "name": "Vile Curse (Costs 3 Actions)"
"lair_actions":
- "desc": "On initiative count 20 (losing initiative ties), the demilich rolls a dice:\
    \ d20|avg|noform (d20). On a result of 11 or higher, the demilich takes a lair\
    \ action to cause one of the following effects. It can't use the same effect two\
    \ rounds in a row."
  "name": ""
- "desc": "- The tomb trembles violently for a moment. Each creature on the floor\
    \ of the tomb must succeed on a DC 19 Dexterity saving throw or be knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone).\
    \  \n- The demilich targets one creature it can see within 60 feet of it. An [antimagic\
    \ field](/3-Mechanics/CLI/spells/antimagic-field.md) fills the space of the target,\
    \ moving with it until initiative count 20 on the next round.  \n- The demilich\
    \ targets any number of creatures it can see within 30 feet of it. No target can\
    \ regain hit points until initiative count 20 on the next round.  "
  "name": ""
"regional_effects":
- "desc": "A demilich's tomb might have any or all of the following effects in place:"
  "name": ""
- "desc": "- The first time a non-evil creature enters the tomb's area, the creature\
    \ takes dice:3d10|text(16) (3d10) necrotic damage.  \n- Monsters in the tomb\
    \ have advantage on saving throws against being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ or [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened), and against\
    \ features that turn undead.  \n- The tomb is warded against the magical travel\
    \ of creatures the demilich hasn't authorized. Such creatures can't teleport into\
    \ or out of the tomb's area or use planar travel to enter or leave it. Effects\
    \ that allow teleportation or planar travel work within the tomb as long as they\
    \ aren't used to leave or enter the tomb's area.  "
  "name": ""
- "desc": "If the demilich is destroyed, these effects fade over the course of 10\
    \ days."
  "name": ""
"source":
- "MM"
- "TftYP"
- "WDMM"
- "GoS"
- "EGW"
- "IDRotF"
- "JttRC"
- "ToFW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Demilich.webp"
```
^statblock

```encounter-table
name: Demilich
creatures:
 - 1: Demilich
```