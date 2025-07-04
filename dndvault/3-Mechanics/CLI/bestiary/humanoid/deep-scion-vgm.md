---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/3
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/humanoid/shapechanger
aliases: ["Deep Scion"]
NoteIcon: monster
BestiaryType: humanoid (shapechanger)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 135
---
# [Deep Scion](3-Mechanics\CLI\bestiary\humanoid/deep-scion-vgm.md)
*Source: Volo's Guide to Monsters p. 135*  

Deep scions began life as people who were stolen from shore or saved from sinking ships and offered a terrible bargain by an undersea power: surrender, body and soul, or drown. Those who submit are subjected to an ancient ritual widespread among evil aquatic creatures. Its methods are painful and the result never certain, but when it works, the magic transforms an air-breathing person into a shapechanger that can take a form that is fully at home beneath the waves.

## Spies from the Sea

A deep scion emerges from the depths in service to its underwater master, which is likely a kraken or some other ancient being of the deep. While wearing the mind and body of the person it once was as a sort of mask, the creature is bent on fulfilling its master's desires. Sometimes a deep scion returns to its former home and a hero's welcome-unexpectedly found alive when all hope was lost. At other times the deep scion takes on a new identity. In any case, it is the deep scion's duty to infiltrate the air-breathing world and report back to its master. When set to its task, a deep scion worms its way into the life of an unsuspecting enemy as a new best friend, an irresistible lover, the perfect candidate for a job, or in some other role that enables the minion to carry out its master's commands.

If you meet a human and there's something fishy about them, they might be a deep scion. Or a crook, or just a fishmonger. Sometimes fish stink is just fish stink.

-Volo

## Cold-Hearted Killers

The training to which a deep scion is subjected rids it of empathy for those whom it spies on. Though one might behave as though infatuated, laugh at the joke of a friend, or appear incensed at some injustice, each of these acts is artificial to the deep scion, a means to an end. It believes that its true form is the shape it takes when it returns to the sea that it thinks of as home. Ironically, however, a deep scion that is killed when in its piscine form is stripped of the magic that robbed it of emotion, leaving behind the corpse of the person the deep scion once was.

```statblock
"name": "Deep Scion (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "shapechanger"
"alignment": "Neutral Evil"
"ac": !!int "11"
"hp": !!int "67"
"hit_dice": "9d8 + 27"
"stats":
- !!int "18"
- !!int "13"
- !!int "16"
- !!int "10"
- !!int "12"
- !!int "14"
"speed": "30 ft. (20 ft. and swim 40 ft. in hybrid form)"
"saves":
  "Charisma": !!int "4"
  "Wisdom": !!int "3"
"skillsaves":
  "Sleight of Hand": !!int "3"
  "Deception": !!int "6"
  "Stealth": !!int "3"
  "Insight": !!int "3"
"senses": "darkvision 120 ft., passive Perception 11"
"languages": "Aquan, Common, Thieves' cant"
"cr": "3"
"traits":
- "desc": "The deep scion can use its action to polymorph into a humanoid-piscine\
    \ hybrid form, or back into its true form. Its statistics, other than its speed,\
    \ are the same in each form. Any equipment it is wearing or carrying isn't transformed.\
    \ The deep scion reverts to its true form if it dies."
  "name": "Shapechanger"
- "desc": "The deep scion can breathe air and water."
  "name": "Amphibious (Hybrid Form Only)"
"actions":
- "desc": "In humanoid form, the deep scion makes two melee attacks. In hybrid form,\
    \ the deep scion makes three attacks: one with its bite and two with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 4|text(8) (1d8 + 4) slashing damage, or dice:1d10 + 4|text(9)\
    \ (1d10 + 4) slashing damage if used with two hands."
  "name": "Battleaxe (Humanoid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d4 + 4|text(6) (1d4 + 4) piercing damage."
  "name": "Bite (Hybrid Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage."
  "name": "Claw (Hybrid Form Only)"
- "desc": "The deep scion emits a terrible scream audible within 300 feet. Creatures\
    \ within 30 feet of the deep scion must succeed on a DC 13 Wisdom saving throw\
    \ or be [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned) until the end\
    \ of the deep scion's next turn. In water, the psychic screech also telepathically\
    \ transmits the deep scion's memories of the last 24 hours to its master, regardless\
    \ of distance, so long as it and its master are in the same body of water."
  "name": "Psychic Screech (Hybrid Form Only; Recharges after a Short or Long Rest)"
"source":
- "VGM"
- "GoS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Deep%20Scion.webp"
```
^statblock

```encounter-table
name: Deep Scion
creatures:
 - 1: Deep Scion
```

## Environment

underwater, coastal