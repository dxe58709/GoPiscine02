/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   iterativefactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 16:18:20 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/27 18:31:55 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb > 0 && nb <= 20 {
		for i := 1; i <= nb; i++ {
			result *= i
		}
	} else {
		return 0
	}
	return result
}
