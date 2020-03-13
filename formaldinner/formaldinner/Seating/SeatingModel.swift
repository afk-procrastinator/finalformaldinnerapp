//
//  SeatingModel.swift
//  formaldinner
//
//  Created by Spencer Michaels on 2/20/20.
//  Copyright © 2020 cate. All rights reserved.
//


import Foundation

struct SeatingModel: Decodable {
    let First: String
    let Last: String
    let Table: Int
    let IsKC: Bool
    let IsW: Bool
    let Value: Int
}

//
//    let main: Main
//struct Main: Codable {
//    // stored properties
//
//}
