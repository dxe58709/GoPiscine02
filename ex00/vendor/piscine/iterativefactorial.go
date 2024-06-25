/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   iterativefactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 16:18:20 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/25 16:52:14 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= nb; i++ {
		if result > 2147483647 {
			return 0
		}
		result *= i
	}
	return result
}
