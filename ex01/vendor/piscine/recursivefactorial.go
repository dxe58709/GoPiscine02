/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   recursivefactorial.go                              :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/25 16:54:12 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/27 18:31:36 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb - 1)
}
