---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/8
- monster/size/large
- monster/type/aberration/shapechanger
aliases: ["Green Slaad"]
NoteIcon: monster
BestiaryType: aberration (shapechanger)
SourceType: Bestiary
BookSource: Monster Manual p. 277
---
# [Green Slaad](3-Mechanics\CLI\bestiary\aberration/green-slaad.md)
*Source: Monster Manual p. 277*  

Green slaadi are surprisingly intelligent and possess innate spellcasting ability. A green slaad can change its shape to appear as a humanoid. If it was born of a humanoid host, the slaad usually adopts its host's form.

At some unpredictable point in its existence, a green slaad unlocks the means to magically, instantly, and permanently transform itself into a gray slaad. Unlocking this knowledge can take years, even decades.

## Slaadi

In the Ever-Changing Chaos of Limbo, bits of forest and meadow, ruined castles, and isolated islands drift through a tumult of fire, water, earth, and wind. The foremost inhabitants of this inhospitable plane are the toad-like slaadi. Slaadi are undisciplined and have no formal hierarchy, although weaker slaadi obey stronger ones under threat of annihilation.

### The Spawning Stone

Long ago, Primus, overlord of the modrons, created a gigantic, geometrically complex stone imbued with the power of law. He then cast it adrift in Limbo, believing that the stone would bring order to the chaos of that plane and halt the spread of chaos to other planes. As the stone's power grew, it became possible for creatures with ordered minds, such as modrons and githzerai, to create enclaves in Limbo. However, Primus's creation had an unforeseen side effect: the chaotic energy absorbed by the stone spawned the horrors that came to be known as slaadi. Sages refer to Primus's massive creation as the Spawning Stone for this reason.

The slaadi wiped out every last modron enclave in Limbo. As creatures of utter chaos, slaadi loathe modrons and attack them on sight. Nonetheless, Primus stands by his creation and either doesn't perceive the slaadi as threats or chooses to ignore them.

### Birth and Transformation

Slaadi have horrific cycles of reproduction. Slaadi reproduce either by implanting humanoid hosts with eggs or by infecting them with a transformative disease called chaos phage. Each color of slaad reproduces or transforms in a different way, with red slaadi spawning blue and green slaadi, and blue slaadi spawning red and green. Each green slaad undergoes a lifelong cycle of transformation into the more powerful gray and death slaadi. With each transformation, the slaad retains its memories.

### Shapechangers

Some slaadi can transform into the humanoid creatures from which they were originally spawned. These slaadi return to the Material Plane to sow discord in the guise of their former selves.

> [!note] Variant: Slaad Control Gems
> 
> As a slaad emerges from the Spawning Stone, the stone magically implants a fragment of itself in the slaad's brain. This fragment takes the form of a magic gem roughly the size and shape of a human child's fist. The gem is the same color as the slaad. Another creature can use magic to draw forth a slaad's gem and use it to subjugate the slaad. The slaad must obey whoever possesses its gem. If a slaad's gem is destroyed, the slaad can no longer be controlled in this way.
> 
> A slaad born from something other than the Spawning Stone has no gem in its brain, but it gains one if it ever comes into contact with the Spawning Stone. Slaadi on Limbo are attracted to the Spawning Stone, so most end up with a gem. A slaad with a control gem in its brain has the following additional trait.
> 
> **Control Gem.** Implanted in the slaad's brain is a magic control gem. The slaad must obey whoever possesses the gem and is immune to being [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) while so controlled.
> 
> Certain spells can be used to acquire the gem. If the slaad fails its saving throw against imprisonment, the spell can transfer the gem to the spellcaster's open hand, instead of imprisoning the slaad. A [wish](/3-Mechanics/CLI/spells/wish.md) spell, if cast in the slaad's presence, can be worded to acquire the gem.
> 
> A [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md) spell cast on the slaad destroys the gem without harming the slaad.
> 
> Someone who is proficient in Wisdom ([Medicine](/3-Mechanics/CLI/rules/skills.md#Medicine)) can remove the gem from an [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated) slaad. Each try requires 1 minute of uninterrupted work and a successful DC 20 Wisdom ([Medicine](/3-Mechanics/CLI/rules/skills.md#Medicine)) check. Each failed attempt deals `dice:4d10|text(22)` (`4d10`) psychic damage to the slaad.
^variant-slaad-control-gems

> [!quote]-  
> 
> Embedded in a slaad's brain is a magic gem. Acquire it, and the slaad is yours to command.


```statblock
"name": "Green Slaad"
"size": "Large"
"type": "aberration"
"subtype": "shapechanger"
"alignment": "Chaotic Neutral"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "127"
"hit_dice": "15d10 + 45"
"stats":
- !!int "18"
- !!int "15"
- !!int "16"
- !!int "11"
- !!int "8"
- !!int "12"
"speed": "30 ft."
"skillsaves":
  "Perception": !!int "2"
  "Arcana": !!int "3"
"damage_resistances": "acid, cold, fire, lightning, thunder"
"senses": "blindsight 30 ft., darkvision 60 ft., passive Perception 12"
"languages": "Slaad, telepathy 60 ft."
"cr": "8"
"traits":
- "desc": "The slaad's innate spellcasting ability is Charisma (spell save DC 12).\
    \ The slaad can innately cast the following spells, requiring no material components:\n\
    \nAt will: [detect magic](/3-Mechanics/CLI/spells/detect-magic.md), [detect\
    \ thoughts](/3-Mechanics/CLI/spells/detect-thoughts.md), [mage hand](/3-Mechanics/CLI/spells/mage-hand.md)\n\
    \n1/day: [fireball](/3-Mechanics/CLI/spells/fireball.md)\n\n2/day each:\
    \ [fear](/3-Mechanics/CLI/spells/fear.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only)"
  "name": "Innate Spellcasting"
- "desc": "The slaad can use its action to polymorph into a Small or Medium humanoid,\
    \ or back into its true form. Its statistics, other than its size, are the same\
    \ in each form. Any equipment it is wearing or carrying isn't transformed. It\
    \ reverts to its true form if it dies."
  "name": "Shapechanger"
- "desc": "The slaad has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The slaad regains 10 hit points at the start of its turn if it has at least\
    \ 1 hit point."
  "name": "Regeneration"
"actions":
- "desc": "The slaad makes three attacks: one with its bite and two with its claws\
    \ or staff. Alternatively, it uses its Hurl Flame twice."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) piercing damage."
  "name": "Bite (Slaad Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) slashing damage."
  "name": "Claw (Slaad Form Only)"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Staff"
- "desc": "Ranged Spell Attack: dice: d20+4 (+4) to hit, range 60 ft., one target.\
    \ Hit: dice:3d6|text(10) (3d6) fire damage. The fire ignites flammable objects\
    \ that aren't being worn or carried."
  "name": "Hurl Flame"
"source":
- "MM"
- "WDMM"
- "IDRotF"
- "DSotDQ"
- "SatO"
- "ToFW"
- "BMT"
- "VEoR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Green%20Slaad.webp"
```
^statblock

```encounter-table
name: Green Slaad
creatures:
 - 1: Green Slaad
```