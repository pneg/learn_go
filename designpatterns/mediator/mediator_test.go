package mediator

import "testing"

func TestMediator(t *testing.T) {
	// 创建机场调度塔台
	airportMediator := &ApproachTower{hasFreeAirstrip: true}
	// 创建C919客机
	c919Airline := NewAirline("C919", airportMediator)
	// 创建米-26重型运输直升机
	m26Helicopter := NewHelicopter("米-26", airportMediator)

	c919Airline.ApproachAirport()   // c919进港降落
	m26Helicopter.ApproachAirport() // 米-26进港等待

	c919Airline.DepartAirport()   // c919飞离，等待的米-26进港降落
	m26Helicopter.DepartAirport() // 最后米-26 飞离
}
