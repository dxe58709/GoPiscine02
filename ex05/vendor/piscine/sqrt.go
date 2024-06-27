/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sqrt.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 16:31:20 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/27 18:21:48 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	for i := uint64(0); i * i <= uint64(nb); i++ {
		if i * i < 0 {
			return -42
		}
		if i * i == uint64(nb) {
			return i
		}
	}
	return 0
}

// func Sqrt(nb int) int {
// 	if nb < 0 {
// 		return 0
// 	}
// 	for i := 0; i * i <= nb; i++ {
// 		if i * i == nb {
// 			return i
// 		}
// 	}
// 	return 0
// }