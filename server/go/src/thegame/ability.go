package main

const (
	HealthRegen = iota
	MaxHealth
	BodyDamage
	BulletSpeed
	BulletPenetration
	BulletDamage
	Reload
	MovementSpeed
	NAbilities
)

var AbilityValues = [NAbilities][9]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8},                            // HealthRegen
	{1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000}, // MaxHealth
	{20, 25, 30, 35, 40, 45, 50, 55, 60},                   // BodyDamage
	{8, 9, 10, 11, 12, 13, 14, 15, 16},                     // BulletSpeed
	{40, 50, 60, 70, 80, 90, 100, 110, 120},                // BulletPenetration
	{12, 15, 18, 21, 24, 27, 30, 33, 36},                   // BulletDamage
	{15, 13, 11, 9, 7, 6, 5, 4, 3},                         // Reload
	{6, 7, 8, 9, 10, 11, 12, 13, 14},                       // MovementSpeed
}